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
    $ /path/to/pure -version --last-command-status $?
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

Enable `PROMPT_SUBST` and single-quote `PROMPT`, so the command runs on
every prompt instead of once when `.zshrc` loads:

```zsh
setopt PROMPT_SUBST
PROMPT='$(/path/to/pure -version --last-command-status $?)'
```

### Bash

```sh
export PS1='$(/path/to/pure -version --last-command-status $?)'
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

## Ports

* Original [sindresorhus/pure](https://github.com/sindresorhus/pure) Zsh implementation ;
* :fish: [pure-fish/pure](https://github.com/pure-fish/pure/) Fish implementation (features rich) ;
* :snake: [edouard-lopez/pure.py](https://github.com/edouard-lopez/pure.py/) Python implementation  (partial features) ;
* :goat: [edouard-lopez/pure.go](https://github.com/edouard-lopez/pure.go/) Golang implementation  (partial features) ;
* :gear: [xcambar/purs](https://github.com/xcambar/purs) Rust implementation ;
* :window: [nickcox/pure-pwsh](https://github.com/nickcox/pure-pwsh/) PowerShell implementation.
* :shell: [krashikiworks/pure-prompt-bash](https://github.com/krashikiworks/pure-prompt-bash) Bash implementation.
