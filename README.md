<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

[![Build Status](https://travis-ci.org/gsamokovarov/jump.svg?branch=main)](https://travis-ci.org/gsamokovarov/jump) [![Go Report Card](https://goreportcard.com/badge/github.com/gsamokovarov/jump)](https://goreportcard.com/report/github.com/gsamokovarov/jump)

**Jump** integrates with your shell and learns about your navigational habits by
keeping track of the directories you visit. It gives you the most visited
directory for the shortest search term you type.

![Demo](./assets/demo.svg)

## Installation

Jump comes in packages for the following platforms.

| Platform | Command |
| --- | --- |
| macOS | `brew install jump` |
| Ubuntu | `wget https://github.com/gsamokovarov/jump/releases/download/v0.40.0/jump_0.40.0_amd64.deb && sudo dpkg -i jump_0.40.0_amd64.deb` |
| Fedora | `wget https://github.com/gsamokovarov/jump/releases/download/v0.40.0/jump-0.40.0-1.x86_64.rpm && sudo rpm -i jump-0.40.0-1.x86_64.rpm` |
| Nix | `nix-env -iA nixpkgs.jump` |
| Go | `go get github.com/gsamokovarov/jump` |

### Integration

Jump needs to be integrated with the shell. For `bash` and `zsh`, the line
below needs to be in `~/.bashrc`, `~/bash_profile` or `~/.zshrc`:

```bash
eval "$(jump shell)"
```

For fish shell, put the line below needs to be in `~/.config/fish/config.fish`:

```
status --is-interactive; and source (jump shell fish | psub)
```

Once integrated, jump will automatically monitor directory changes and start
building an internal database.


### But `j` is not my favourite letter!

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
of the directory name is all that **jump** needs to find it.

### Regular jump

The default search behavior of **jump** is to fuzzy match the
directory name of a score. The match is case insensitive.

If you visit the directory `/Users/genadi/Development/rails/web-console` often,
you can jump to it by:

```bash
$ j wc      # or...
$ j webc    # or...
$ j console # or...
$ j b-c     # or...
```

Using jump is all about saving key strokes. However, if you made the effort to
type a directory base name exactly, **jump** will try to find the exact match,
rather than fuzzy search.

```bash
$ j web-console
$ pwd
/Users/genadi/Development/rails/web-console
```

### Deep jump

Given the following directories:

```bash
/Users/genadi/Development/society/website
/Users/genadi/Development/chaos/website
```

Typing `j site` matches only the base names of the directories. The base name
of `/Users/genadi/Development/society/website` is `website`, the same as the
other absolute path above. The jump above will land on the most scrored path,
which is the `society` one, however what if we wanted to land on the `chaos`
website?

```bash
$ j ch site
$ pwd
/Users/genadi/Development/chaos/website
```

This instructs **jump** to look for a `site` match inside that is preceded by a
`ch` match in the parent directory. The search is normalized only on the last
two parts of the target paths. This will ensure a better match, because of the
shorter path to fuzzy match on.

There are no depth limitations though and a jump to
`/Users/genadi/Development/society/website` can look like:

```bash
$ j dev soc web
$ pwd
/Users/genadi/Development/society/website
```

In fact, every space passed to `j` is converted to an OS separator. The last
search term can be expressed as:

```bash
$ j dev/soc/web
$ pwd
/Users/genadi/Development/society/website
```

## Reverse jump

Bad jumps happen. Sometimes we're looking for a directory the isn't the most
scored one at the moment. Imagine the following jump database:

```bash
/Users/genadi/Development/society/website
/Users/genadi/Development/chaos/website
/Users/genadi/Development/hack/website
```

Typing `j web` would lead to:

```bash
$ j web
$ pwd
/Users/genadi/Development/society/website
```

If we didn't expect this result, instead of another search term, typing **j**
without any arguments will instruct **jump** to go the second best match.

```bash
$ j
$ pwd
/Users/genadi/Development/chaos/website
```

### Case sensitive jump

To trigger a case-sensitive search, use a term that has a capital letter.

```bash
$ j Dev
$ pwd
/Users/genadi/Development
```

The jump will resolve to `/Users/genadi/Development` even if there is
`/Users/genadi/Development/dev-tools` that scores better.

## Is it like autojump or z?

Yes, it is! You can import your datafile from `autojump` or `z` with:

```bash
$ jump import
```

This will try `z` first then `autojump`, so you can even combine all the
entries from both tools.

The command is safe to run on pre-existing jump database, because if an entry
exist in jump already, it won't be imported and it's score will remain
unchanged. You can be explicit and choose to import `autojump` or `z` with:

```bash
$ jump import autojump
$ jump import z
```

If you want to know more about the difference between Jump, z, and autojump,
check-out this Twitter [conversation].

## Thanks! 🙌

Thank you for stopping by and showing your interest in Jump!

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[conversation]: https://twitter.com/hkdobrev/status/838398833419767808
