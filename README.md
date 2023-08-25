# pure.go

> Exploring Golang by re-implementing pure prompt

## Usage

```sh
❯ go run ./cmd/cli.go -version --last-command-status $status
```

or the built binary

```sh
❯ make build
❯ ./pure --last-command-status $status
```

## Build

```sh
❯ make build
```

## Todo

* [x] create a `pure` package
  * [x] add `pure` package to `go.mod`
  * [x] add `pure` package to `go.sum`
* [x] print prompt `❯` with `pure` package
* [x] test pure package output (with [`be`][be])
* [x] create a CLI (with [`clîr`][cli])
* [ ] add current working directory to prompt
* [ ] add `go version` to prompt when `go.sum` exists

[be]: https://github.com/carlmjohnson/be
[cli]: https://github.com/leaanthony/clir
