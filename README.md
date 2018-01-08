<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

**Jump** integrates with the shell and learns about your navigational habits by
keeping track of the directories you visit. It strives to give you the best
directory for the shortest search term.

![demo](https://raw.githubusercontent.com/gsamokovarov/jump/master/assets/demo.gif)

### Integration

Jump needs to be integrated with the shell. For `bash` and `zsh`, the the line
below in needs to be in `~/.bashrc`, `~/bash_profile` or `~/.zshrc`:

    eval "$(jump shell)"

For fish shell, put the line below needs to be in `~/.config/fish/config.fish`:

    status --is-interactive; and . (jump shell | psub)

Once integrated, jump will automatically directory changes and start
building an internal database.


### But `j` is not my favourite letter!

<img align="right" src="https://github.com/gsamokovarov/jump/raw/master/assets/logo-light.png">

This is fine, you can bind jump to `z`, with this:

```bash
eval "$(jump shell --bind=z)"
```

And now, you can use `jump` like `z dir` and it would just work! This is only
an example, you can bind it to _anything_. If you are one of those persons that
likes to type a lot with their fingers, you can do:

```bash
eval "$(jump shell --bind=goto)"
```

Voila! `goto dir` becomes a thing. The possibilities are endless!

## Usage

Once integrated, **jump** introduces the **j** helper. It accepts only search
terms and as a design goal there are no arguments to **j**. Whatever you give
it, it's treated as search term.

**Jump** uses fuzzy matching to find the desired directory to jump to. This
means that your search terms are patterns that match the desired directory
approximately rather than exactly. Typing **2** to **5** consecutive characters
of the base directory names is all that **jump** needs to find it.

### Regular jump

The default search behavior of **jump** is to case insensitively match only the
base directory path of the scored directories. This is because absolute paths
are long and short search terms can fuzzy match them easily, lead to bad
matches.

If you visit the directory `/Users/genadi/Development/rails/web-console` often,
you can jump to it by:

    $ j wc      # or...
    $ j webc    # or...
    $ j console # or...
    $ j b-c     # or...

Of course, `web-console` can be typed directly as a search term:

    $ j web-console
    $ pwd
    /Users/genadi/Development/rails/web-console

Using jump is all about saving key strokes. However, if you made the effort to
type a directory base name exactly, **jump** will try to find the exact match,
rather than fuzzy search.

### Deep jump

Given the following directories:

    /Users/genadi/Development/society/website
    /Users/genadi/Development/chaos/website

Can you be sure where `j web` will lead you? You can hint jump where you want
to go.  To ensure a match of `/Users/genadi/Development/chaos/website`, use the
search term:

    $ j ch web
    $ pwd
    /Users/genadi/Development/chaos/website

This instructs **jump** to look for a `web` match inside that is preceded by a
`ch` match in the parent directory.  The search is normalized only on the last
two parts of the target paths. This will ensure a better match, because of the
shorter path to fuzzy match on.

There are no depth limitations though and a jump to
`/Users/genadi/Development/society/website` can look like:

    $ j dev soc web
    $ pwd
    /Users/genadi/Development/society/website

In fact, every space passed to `j` is converted to an OS separator. The search
term above can be expressed as:

    $ j dev/soc/web
    $ pwd
    /Users/genadi/Development/society/website

## Reverse jump

Sometimes bad jumps happen. Maybe the search has a better scored directory
already. If we want to jump to `/Users/genadi/Development/hack/website` and we
have the following entries in the database:

    /Users/genadi/Development/society/website
    /Users/genadi/Development/chaos/website
    /Users/genadi/Development/hack/website

Typing `j web` would lead to:

    $ j web
    $ pwd
    /Users/genadi/Development/society/website

Instead of typing another search term, typing **j** without a search term will
instruct **jump** to the second best, third best and so on matches.

    $ j
    $ pwd
    /Users/genadi/Development/chaos/website

    $ j
    $ pwd
    /Users/genadi/Development/hack/website

### Case sensitive jump

To trigger a case-sensitive search, use a term that has a capital letter.

    $ j Dev
    $ pwd
    /Users/genadi/Development

The jump will resolve to `/Users/genadi/Development` even if there is
`/Users/genadi/Development/dev-tools` that scores better.


## Installation

Jump comes in packages for macOS through homebrew and linux.

## macOS

```bash
brew install jump
```

### Ubuntu/Debian

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.13.0/jump_0.13.0_amd64.deb
sudo dpkg -i jump_0.13.0_amd64.deb
```

### Red Hat/Fedora

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.13.0/jump-0.13.0-1.x86_64.rpm
sudo rpm -i jump-0.13.0-1.x86_64.rpm
```

### Go

If you have the Go toolchain installed, you can install it through:

```bash
go get github.com/gsamokovarov/jump
```

You can also build jump by yourself. Or hack on it, you know, if you like Go
and UNIX stuff. 💻

## Build

[![Build Status](https://travis-ci.org/gsamokovarov/jump.svg?branch=master)](https://travis-ci.org/gsamokovarov/jump) [![Go Report Card](https://goreportcard.com/badge/github.com/gsamokovarov/jump)](https://goreportcard.com/report/github.com/gsamokovarov/jump)

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[conversation]: https://twitter.com/hkdobrev/status/838398833419767808
