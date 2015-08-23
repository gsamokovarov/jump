# jump

Like [autojump] or [z], but fuzzy.

## Overview

Jump keeps track of the visited directories in bash, zsh and fish. After a
while, the `j` command can be used to jump to one of your most recently visited
directories. Say, you visit `/Projects/rails` often. Typing `j rls` will take
you straight to it.

Yes, it is fuzzy.

## Installation

Right now, you can only install it through source:

```shell
git clone https://github.com/gsamokovarov/jump
cd jump
make
mv jump ~/bin # Or /usr/local/bin, if ~/bin isn't in $PATH.
```

## Shell

To get full advantage of jump, you wanna integrate it with you shell. The
integration will track the directories entered through `cd`, `pushdir`,
`popdir` and the likes. It will also provide the `j` command, so you can jump
around easier.

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

## Progress

This is still a work in progress. The concept is there, but a lot of things
will change.

Hope you find jump useful.

[autojump]: https://github.com/wting/autojump
[z]: https://github.com/rupa/z
