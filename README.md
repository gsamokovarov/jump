# jump

A fuzzy quick-directory jumper. Kinda like [autojump] or [z], but  fuzzy.

Jump works its magic by keeping track of the directories you visit. It scores
them to give you the best match for tour input. When integrated with your
shell, the `j` function is available. It let's you jump across directories with
ease.

If you visit `/Users/bob/Projects/moneymaker` often, type `j mk` and jump
straight to it. Gone are the days of manual aliases for frequent project
directories.

## Usage

The default search behavior of `jump` is to match case insensitively only the
base path (only the last directory name of the full absolute path) of the
scored directories.

Why we wanna do that? Because directory names are long and short terms can
fuzzy match them easily. One of the goals of `jump` is to type less. This
assumption helps with that.

### Case-sensitive Search

To trigger a case-sensitive search, use a term that has different case in one
of the letters.

For example:

```bash
j Dev
```

Will jump to `/Users/foo/Development` instead of
`/Users/foo/Development/dev-tools` even if `dev-tools` has scored better.

### Deeper Search

The first `jump` normalized all search terms to the base names the saved
directories. Why? Because directory names are long and short terms can fuzzy
match them easily. Most of the times we actually want that. But only most of
the times. :-)

Say you have a lot of client specific directory with projects inside of them.

```bash
society/
├── artwork
├── interview
└── website

chaos/
└── website
```

When you time `web`, which `website` should you jump to? Currently, it will be
the one with the higher score, say `society/website`. If you wanted to go to
`chaos/website`, you had no way to trigger it. Well, now you can:

```bash
j ch/web
```

The term above will match `/Users/foo/Development/chaos/website`. The search is
normalized only on the last two parts of the target paths. This will again
ensure you better match, because the path gets shorter.

You can put as many separators as you want in your term.

```bash
j dev/ch/web
```

The term above will match the last three directories of the path.

## Installation

### OSX

```shell
brew tap gsamokovarov/jump
brew install jump
```

### Debian

```shell
wget https://github.com/gsamokovarov/jump/releases/download/v0.2.0/jump_0.2.0_amd64.deb
sudo dpkg -i jump_0.2.0_amd64.deb
```

### Red Hat

```shell
wget https://github.com/gsamokovarov/jump/releases/download/v0.2.0/jump-0.2.0-1.x86_64.rpm
sudo rpm -i jump-0.2.0-1.x86_64.rpm
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
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[issue tracker]: https://github.com/gsamokovarov/jump/issues
[MIT license]: https://github.com/gsamokovarov/jump/blob/master/LICENSE.txt
