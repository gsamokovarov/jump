# JUMP 🏃

<img align="right" src="https://github.com/gsamokovarov/jump/raw/master/assets/logo-light.png">

Jump helps you navigate your file system faster by learning your
habits.

Say you visit `/Users/genadi/Development/web-console` a lot. Jump can
get you there with `j wc` or `j web` or `j webc`. You name it,
loosely, and jump will figure it out for you.

This comes with zero configuration! Install jump, integrate it to your
shell and let it learn your habits for a while – cd to your
directories like you always do. After a while, jump would know how to
get you when you type `j somewhere` or just `j some`.

Maybe you made a typo like `j ssome`? No problem, jump uses fuzzy
searching, so you can type tiny search terms (mostly 2 or 3 characters
are enough) and be tolerated even when you have typos.

## Usage

To get the most out of jump, you have to integrate it with your shell. The
integration gives you the `j` shell function and the automatic tracking and
scoring.

### Shell Integration

Put the line below in `~/.bashrc`,  `~/bash_profile` or `.zshrc` for
zshell:

```bash
eval "$(jump shell)"
```

Put the line below in `~/.config/fish/config.fish` for fish shell:

```fish
status --is-interactive; and . (jump shell | psub)
```

Once the integration is done, work like you always do. In a while you
can just `j` to your projects from everywhere. See [`man jump`][man]
for more usage patterns.

But hey, `j` is not my favourite word, you may say. This is fine,
you can bind jump to `z`, with this:

```bash
eval "$(jump shell --bind=z)"
```

And now, you can use `jump` like `z dir` and it would work! This is
just an example, you can bind it to _anything_. If you are one of
those persons that likes to type, with their fingers:

```bash
eval "$(jump shell --bind=goto)"
```

Voila! `goto dir` becomes a thing. The possibilities are endless!

## Installation

### OSX

```bash
brew install jump
```

### Debian

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.10.0/jump_0.10.0_amd64.deb
sudo dpkg -i jump_0.10.0_amd64.deb
```

### Red Hat

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.10.0/jump-0.10.0-1.x86_64.rpm
sudo rpm -i jump-0.10.0-1.x86_64.rpm
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

## Issues

If you find any problems with jump, please, consider reporting them to the
[issue tracker].

## License

Jump is licensed under the [MIT license].

Hope you find jump useful! :sparkles:

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[issue tracker]: https://github.com/gsamokovarov/jump/issues
[MIT license]: https://github.com/gsamokovarov/jump/blob/master/LICENSE.txt
