git-substitute
==============

the stupid search and replacer

[![Build Status](https://img.shields.io/travis/nicknovitski/git-substitute.svg?style=flat-square)][travis]

[travis]: https://travis-ci.org/nicknovitski/git-substitute

## Installation

If you've already installed Go, and `$GOPATH/bin` is in your `$PATH`, then just:
```shell
$ go get github.com/nicknovitski/git-substitute
```

Alternately, you can download a prebuilt binary from the [latest release](https://github.com/nicknovitski/git-substitute/releases/latest).

Currently the only prebuilt binary is linux/amd64.  Darwin and 386 coming Real Soon Now.

## Usage

Pass a search pattern and a replacement string to replace text matching
the former with the latter.

```shell
$ git substitute foo bar # "foo" -> "bar"
$ git substitute people?s persons # "people" & "peoples" -> "persons"
```

Pass one or more paths to restrict the substitution.
```shell
$ git substitute Command Demand bin doc # "Command" -> "Demand", but only in bin/ and doc/
```

Remember that your shell will get confused by parens and backslashes
unless you wrap the patterns in quotes.
```shell
git substitute '\bUser\.find_by_name\((.*)\)' 'User.where(name: \1).first'
```

Finally, before you dive in, may I recommend an alias to `git sub`?  It's much less typing.
```shell
git config --global alias.sub substitute
```

I would have just named the command `sub`, but:

1. The other git commands are all complete English words.
2. SEO.

## Acknowledgments

I can't remember when I first read [jason meredith's grep-l-pipe-xargs-sed-i-s method]
(http://blog.jasonmeridth.com/posts/use-git-grep-to-replace-strings-in-files-in-your-git-repository/),
but I remember many times since then when it's enabled me to trivially accomplish enormous things.
