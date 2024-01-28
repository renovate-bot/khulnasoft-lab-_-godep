# github.com/khulnasoft-lab/godep

This repository hosts the source code of the [pkg.go.dev](https://pkg.go.dev) website,
and [`godep`](https://pkg.go.dev/github.com/khulnasoft-lab/godep/cmd/godep), a documentation
server program.

[![Go Reference](https://pkg.go.dev/badge/github.com/khulnasoft-lab/godep.svg)](https://pkg.go.dev/github.com/khulnasoft-lab/godep)

## pkg.go.dev: a site for discovering Go packages

Pkg.go.dev is a website for discovering and evaluating Go packages and modules.

You can check it out at [https://pkg.go.dev](https://pkg.go.dev).

## godep: a documentation server

`godep` program extracts and generates documentation for Go projects.

Example usage:

```
$ go install github.com/khulnasoft-lab/godep/cmd/godep@latest
$ cd myproject
$ godep -open .
```

For more information, see the [godep documentation](https://pkg.go.dev/github.com/khulnasoft-lab/godep/cmd/godep).

## Requirements

Godep requires Go 1.19 to run.
The last commit that works with Go 1.18 is
9ffe8b928e4fbd3ff7dcf984254629a47f8b6e63.
The last commit that works with Go 1.17 is
4d836c6a652cde92f433967680dfd6171a91ec12.

## Issues

If you want to report a bug or have a feature suggestion, please first check
the [known issues](https://github.com/golang/go/labels/godep) to see if your
issue is already being discussed. If an issue does not already exist, feel free
to [file an issue](https://golang.org/s/godep-feedback).

For answers to frequently asked questions, see [pkg.go.dev/about](https://pkg.go.dev/about).

You can also chat with us on the
[#godep Slack channel](https://gophers.slack.com/archives/C0166L4QGJV) on the
[Gophers Slack](https://invite.slack.golangbridge.org).

## Contributing

We would love your help!

Our canonical Git repository is located at
[go.googlesource.com/godep](https://github.com/khulnasoft-lab/godep).
There is a mirror of the repository at
[github.com/golang/godep](https://github.com/golang/godep).

To contribute, please read our [contributing guide](CONTRIBUTING.md).

## License

Unless otherwise noted, the Go source files are distributed under the BSD-style
license found in the [LICENSE](LICENSE) file.

## Links

- [Homepage](https://pkg.go.dev)
- [Feedback](https://golang.org/s/godep-feedback)
- [Issue Tracker](https://golang.org/s/godep-issues)
