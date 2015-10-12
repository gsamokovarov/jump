# jump

A quick and fuzzy directory jumper. Kinda like [autojump] or [z], but fuzzy.

Jump works its magic by keeping track of the directories you visit. It scores
them to give you the best match for tour input. When integrated with your
shell, the `j` function is available. It let's you jump across directories with
ease.

If you visit `/Users/bob/Projects/website` often, type `j ws` and jump
straight to it. Gone are the days of manual aliases for frequent project
directories.

## Usage

See [`man jump`][man].

## Installation

### OSX

```bash
brew tap gsamokovarov/jump
brew install jump
```

### Debian

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.3.0/jump_0.3.0_amd64.deb
sudo dpkg -i jump_0.3.0_amd64.deb
```

### Red Hat

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.3.0/jump-0.3.0-1.x86_64.rpm
sudo rpm -i jump-0.3.0-1.x86_64.rpm
```

### Source

You need a working [Go workspace].

```bash
go get github.com/gsamokovarov/jump
git clone https://github.com/gsamokovarov/jump
cd jump
make
mv jump ~/bin # Or /usr/local/bin, if ~/bin isn't in $PATH.
```

## Shell

Jump supports bash, zsh and fish out of the box. If your favourite shell isn't
in the list below, give a heads up in the [issue tracker].

To get the most out of jump, you have to integrate it with your shell. The
integration gives you the `j` shell function and the automatic tracking and
scoring.

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

## Issues

If you find any problems with jump, please, consider reporting them to the
[issue tracker].

## License

Jump is licensed under the [MIT license].

Hope you find jump useful! :sparkles:

[autojump]: https://github.com/wting/autojump
[z]: https://github.com/rupa/z
[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[issue tracker]: https://github.com/gsamokovarov/jump/issues
[MIT license]: https://github.com/gsamokovarov/jump/blob/master/LICENSE.txt
