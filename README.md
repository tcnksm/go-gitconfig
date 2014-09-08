go-gitconfig [![GoDoc](https://godoc.org/github.com/tcnksm/go-gitconfig?status.svg)](https://godoc.org/github.com/tcnksm/go-gitconfig) [![Build Status](http://drone.io/github.com/tcnksm/go-gitconfig/status.png)](https://drone.io/github.com/tcnksm/go-gitconfig/latest) [![Coverage Status](http://coveralls.io/repos/tcnksm/go-gitconfig/badge.png)](https://coveralls.io/r/tcnksm/go-gitconfig)
====

Use `gitconfig` values in Golang.

## Usage

If you want to use git user name:

```go
username, err := gitconfig.Global("user.name")
```

Or, if you want to extract origin url of current project:

```go
url, err := gitconfig.Local("remote.origin.url")
```

See more document at [https://godoc.org/github.com/tcnksm/go-gitconfig](https://godoc.org/github.com/tcnksm/go-gitconfig). 

## Install

To install, use `go get`:

```bash
$ go get -d github.com/tcnksm/go-gitconfig
```

## Contribution

1. Fork ([https://github.com/tcnksm/go-gitconfig/fork](https://github.com/tcnksm/go-gitconfig/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## LICENCE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/tcnksm/go-gitconfig/blob/master/LICENCE)

## Author

[tcnksm](https://github.com/tcnksm)
