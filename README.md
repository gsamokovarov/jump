# jump

A fuzzy quick-directory jumper. Kinda like [autojump] or [z], but you know, fuzzy.

Jump works its magic by keeping track of the directories you visit, scoring
them, and then trying to find the best directory for your input.

Jump integrates with your shell and gives you the `j` function. It let's you
jump across directories with ease. Gone are the manual aliases for frequent
project directories.

If you visit `/Users/bob/Projects/moneymaker` often, type `j mk` or `j mmk` and
be done. Fuzzy matching makes everything better. :-)

## Installation

### OSX

```shell
brew tap gsamokovarov/jump
brew install jump --HEAD
```

### Source

On Linux and other UNIX-like systems, that can compile Go code, you can install
jump from source.

You need a working [Go workspace] for the compilation.

```bash
go get github.com/gsamokovarov/jump
git clone https://github.com/gsamokovarov/jump
cd jump
make
mv jump ~/bin # Or /usr/local/bin, if ~/bin isn't in $PATH.
```

## Shell

Jump supports the most popular shells out there, right out of the box. If your
favourite shell isn't in the list below, why not contribute an integration for
it?

The `j` shell function and the automatic directories score comes from those
integrations, so make sure you have set them up.

### bash

Put the line below in `~/.bashrc` or `~/bash_profile`:

```bash
eval "$(jump shell bash)"
```

### zsh

Put the line below in `~/.zshrc`:

```zsh
eval "$(jump shell zsh)"
```

### fish

Put the line below in `~/.config/fish/config.fish`:

```fish
status --is-interactive; and . (jump shell fish | psub)
```

## P.S.

This is still a work in progress, but its getting close to `0.1`. Jump _should_
behave reasonable, and if it doesn't for you, please consider leaving an issue.
It will help me make it better.

Hope you find jump useful! :sparkles:

[autojump]: https://github.com/wting/autojump
[z]: https://github.com/rupa/z
[Go workspace]: https://golang.org/doc/code.html#Workspaces
