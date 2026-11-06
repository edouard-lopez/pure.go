# From Fish to Go

How to port a [pure-fish/pure] feature to pure.go.
Read [How Pure works][how-it-works] first: it explains the Fish pipeline.
The [gap matrix](README.md#gap-matrix) lists what is left to port.

## Mental model

| Fish Pure | pure.go |
| :-------- | :------ |
| Fish calls `fish_prompt` before each prompt. | The shell runs the `pure` binary before each prompt. |
| A segment is a function `_pure_prompt_<name>`. | A segment is a package `internal/segments/<name>` with a `Get()` function. |
| Segments read Fish variables (`$pipestatus`, `$CMD_DURATION`). | The binary cannot see them. The shell passes them as flags or env vars. |
| Settings are `pure_*` variables. | Settings are `PURE_*` env vars, read with `config.Get(name, fallback)`. |
| `set_color` prints colours. | `internal/colorize` wraps `gookit/color`. |
| Fishtape tests per function. | `go test` per package, plus [agnostic-tests] for the whole binary. |

## Idioms

| Fish | Go |
| :--- | :- |
| `set --local name value` | `name := value` |
| `set --query var; and test "$var" = true` | `config.Get("PURE_VAR", "false") == "true"` |
| `test -n "$x"` | `x != ""` |
| `(command git status --porcelain)` | `exec.Command("git", "status", "--porcelain").Output()` |
| `string replace --regex 'a' 'b' $s` | `regexp.MustCompile("a").ReplaceAllString(s, "b")` |
| `string split '/' $p \| tail -n 2` | `parts := strings.Split(p, "/"); parts[len(parts)-2:]` |
| `string length` on visible text | strip ANSI first, then `utf8.RuneCountInString` (or a width lib for emoji) |
| `math "$ms / 1000"` | `time.Duration(ms) * time.Millisecond` |
| `return 1` (failure) | return `""` for "hide segment", `error` for real failures |
| `echo` to build output | return a `string`; only `cmd/cli.go` prints |

## Recipe for one segment (TDD)

1. **Pick** one row of the gap matrix, in phase order.
2. **Read** the Fish function and its test: `functions/_pure_prompt_<name>.fish`, `tests/_pure_prompt_<name>.test.fish`.
3. **Red**: copy each Fishtape case into a table test in `internal/segments/<name>/<name>_test.go`. Run `make test` and see it fail.
4. **Green**: write the smallest `Get()` that passes. Read inputs through `config.Get` and `os.Getenv`.
5. **Wire**: add a field to `prompt.Prompt` and to `PromptLayout`. Keep the Fish order.
6. **Spec**: add or update a black-box spec in [agnostic-tests] so every port gets the same check.
7. **Refactor**, then update the gap matrix row.

Keep `Get()` pure where you can: pass inputs as arguments, keep `os` and `exec` calls at the edge.
That makes tests fast and avoids `os.Chdir` in tests.

## Shell state contract

| Needed by | Fish source | Proposed input |
| :-------- | :---------- | :------------- |
| symbol, exit statuses | `$pipestatus` | `--last-command-status N` (exists), later `--pipestatus "0 1"` |
| command duration | `$CMD_DURATION` | `PURE_CMD_DURATION` (ms) |
| folder shortening | `$COLUMNS` | `COLUMNS` |
| jobs | `jobs -p` | `--jobs N` |
| vi mode | `$fish_bind_mode` | `--bind-mode insert` |
| transient prompt | `--final-rendering` | `--final-rendering` |

Agree on each name in [agnostic-tests] before you code it, so Python and Rust use the same flag.

## Pitfalls

* **Speed**: each `git` call costs time. Fish uses one call per indicator. Batch with `git status --porcelain=v2 --branch` when you add ahead/behind.
* **Detached HEAD**: `git name-rev` is slow in large or `jj` repos (see pure-fish/pure#422).
* **Width**: measure the visible width without ANSI codes, like `_pure_string_width`.
* **Colours**: Fish settings point to a base colour by name (`pure_color_git_branch=pure_color_mute`). Resolve that indirection in `colorize`.
* **Never print** from a segment. Text on stdout ends up in the prompt.

[pure-fish/pure]: https://github.com/pure-fish/pure
[how-it-works]: https://pure-fish.github.io/pure/doc/#how-it-works
[agnostic-tests]: https://github.com/pure-fish/agnostic-tests
