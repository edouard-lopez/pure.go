# pure.go 🐐 [![CI](https://github.com/edouard-lopez/pure.go/actions/workflows/ci.yml/badge.svg)](https://github.com/edouard-lopez/pure.go/actions/workflows/ci.yml)

> Exploring [Golang][go] by re-implementing [pure][pure] prompt.
>
> :information_source: This is a really basic and naive implementation.

### Preview

![preview pure.go](./preview.png)

## Installation

1. Download the [release] binary for your OS/Arch
2. Allow execution

    ```sh
    chmod u+x /path/to/pure
    ```

3. Try to execute manually

    ```sh
    $ /path/to/pure -version --last-command-status $status
    # should print something like:
    ~/projects/contributions/pure.go 🐐go1.20.3
    0 ❯
    ```

4. Edit your [Fish](#fish), [Zsh](#zsh) or [Bash](#bash) prompt (see below).

## Configuring Shell Prompt

### Fish

Edit `$__fish_config_dir/functions/fish_prompt.fish` with:

```fish
function fish_prompt
    /path/to/pure -version --last-command-status $status
end
```

### Zsh

```zsh
PROMPT=$(/path/to/pure -version --last-command-status $status)
```

### Bash

```sh
export PS1='$(/path/to/pure -version --last-command-status $status)'
```

## Development

### Usage

```sh
❯ go run ./cmd/cli.go -version --last-command-status $status
```

or with the built binary

```sh
❯ make build
❯ ./pure --last-command-status $status
# or
❯ make demo
```

### Install

```sh
❯ make install
```

### Test

```sh
❯ make test
❯ make test-agnostic  # cross-implementation suite from pure-fish/agnostic-tests (needs bats)
```

### Configuration

Set `PURE_*` environment variables, named like Fish's `pure_*` variables in upper case:

| Variable | Default | Description |
| :------- | :------ | :---------- |
| `PURE_SYMBOL_PROMPT` | `❯` | Prompt symbol. |
| `PURE_SYMBOL_GIT_DIRTY` | `*` | Shown after the branch when the repository has changes. |

## Build

Create a `pure` binary in the project's root directory.

```sh
❯ make build
❯ ./pure --last-command-status $status
```

## Todo

* [x] create a `pure` package
  * [x] add `pure` package to `go.mod`
  * [x] add `pure` package to `go.sum`
* [x] print prompt `❯` with `pure` package
* [x] test pure package output (with [`testify`][test])
* [x] create a CLI (with [`clîr`][cli])
* [x] add current working directory to prompt
* [x] add `go version` to prompt when `go.sum` exists
* [x] configure Goreleaser
* [x] add color to output
* [x] Add Git segments
  * [x] add Git branch segment
  * [x] add Git dirty segment
* [x] start the prompt with a blank line
* [x] pass [pure-fish/agnostic-tests](https://github.com/pure-fish/agnostic-tests) in CI

[test]: https://github.com/stretchr/testify/
[cli]: https://github.com/leaanthony/clir
[release]: https://github.com/edouard-lopez/pure.go/releases/
[pure]: https://github.com/pure-fish/pure/
[go]: go.dev/

## Gap matrix

Parity with the reference [pure-fish/pure](https://github.com/pure-fish/pure) implementation.
See [How Pure works](https://pure-fish.github.io/pure/doc/#how-it-works) for the Fish internals, and the [From Fish to Go](https://github.com/edouard-lopez/pure.go/wiki/From-Fish-to-Go) wiki page for the porting guide.

✅ done · 🟡 partial · ❌ missing. Phases match the repository milestones.

| Phase | Feature | Fish setting | Status |
| :---- | :------ | :----------- | :----- |
| P0 Foundation | CLI contract `--last-command-status N` | — | ✅ |
| P0 Foundation | Settings from `PURE_*` environment variables | `pure_*` | 🟡 2 settings |
| P0 Foundation | Nine base colours, configurable | `pure_color_*` | 🟡 fixed |
| P0 Foundation | Agnostic tests in CI | — | ✅ |
| P1 Core prompt | Prompt symbol `❯` | `pure_symbol_prompt` | ✅ |
| P1 Core prompt | Symbol red after a failure, magenta after a success | `pure_color_prompt_on_*` | ✅ |
| P1 Core prompt | Separate error symbol | `pure_separate_prompt_on_error` | ❌ |
| P1 Core prompt | Exit statuses list, signal names | `pure_show_exit_status` | ❌ |
| P1 Core prompt | Current folder with `~` | `pure_color_current_directory` | 🟡 `~` only for `/home/$USER` |
| P1 Core prompt | Shorten and truncate folder | `pure_shorten_*`, `pure_truncate_*` | ❌ |
| P1 Core prompt | Blank line before the prompt | — | ✅ |
| P1 Core prompt | Single-line (compact) prompt | `pure_enable_single_line_prompt` | ❌ |
| P1 Core prompt | Folder first or last | `pure_begin_prompt_with_current_directory` | 🟡 folder first only |
| P2 Git | Git branch (and detached HEAD) | `pure_color_git_branch` | ✅ |
| P2 Git | Git dirty `*` | `pure_symbol_git_dirty` | ✅ |
| P2 Git | Git stash `≡` | `pure_symbol_git_stash` | ❌ |
| P2 Git | Git ahead `⇡` / behind `⇣` | `pure_symbol_git_unpushed_commits` | ❌ |
| P2 Git | Numbered git indicators | `pure_show_numbered_git_indicator` | ❌ |
| P2 Git | Disable git | `pure_enable_git` | ❌ |
| P3 Context segments | Command duration | `pure_threshold_command_duration` | ❌ |
| P3 Context segments | `user@host` over SSH | `pure_symbol_ssh_prefix` | ❌ |
| P3 Context segments | Container detection | `pure_enable_container_detection` | ❌ |
| P3 Context segments | Python virtualenv / conda | `pure_enable_virtualenv` | ❌ |
| P3 Context segments | AWS profile | `pure_enable_aws_profile` | ❌ |
| P3 Context segments | Kubernetes context/namespace | `pure_enable_k8s` | ❌ |
| P3 Context segments | Nix dev shell | `pure_enable_nixdevshell` | ❌ |
| P3 Context segments | Background jobs | `pure_show_jobs` | ❌ |
| P3 Context segments | System time | `pure_show_system_time` | ❌ |
| P3 Context segments | Root prefix `#` | `pure_show_prefix_root_prompt` | ❌ |
| P4 Shell integration | Vi mode `❮` / `[N]` | `pure_reverse_prompt_symbol_in_vimode` | ❌ |
| P4 Shell integration | Private mode prefix `!` | `pure_show_prefix_private_prompt` | ❌ |
| P4 Shell integration | Transient prompt | `fish_transient_prompt` | ❌ |
| P4 Shell integration | Window title | `pure_symbol_title_bar_separator` | ❌ |
| P4 Shell integration | iTerm2 prompt mark | — | ❌ |
| P5 Release | Check for new release | `pure_check_for_new_release` | ❌ |
| Extra | Go version `🐐` | — | ➕ not in Fish Pure |

## Ports

* Original [sindresorhus/pure](https://github.com/sindresorhus/pure) Zsh implementation ;
* :fish: [pure-fish/pure](https://github.com/pure-fish/pure/) Fish implementation (features rich) ;
* :snake: [edouard-lopez/pure.py](https://github.com/edouard-lopez/pure.py/) Python implementation  (partial features) ;
* :goat: [edouard-lopez/pure.go](https://github.com/edouard-lopez/pure.go/) Golang implementation  (partial features) ;
* :gear: [xcambar/purs](https://github.com/xcambar/purs) Rust implementation ;
* :window: [nickcox/pure-pwsh](https://github.com/nickcox/pure-pwsh/) PowerShell implementation.
* :shell: [krashikiworks/pure-prompt-bash](https://github.com/krashikiworks/pure-prompt-bash) Bash implementation.
