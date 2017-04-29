<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

<img align="right" src="https://github.com/gsamokovarov/jump/raw/master/assets/logo-light.png">

Jump helps you navigate your file system faster by learning your
habits.

Say you visit `/Users/genadi/Development/web-console` a lot. Jump can
get you there with `j wc` or `j web` or `j webc`. You name it,
loosely, and jump will figure it out for you.

This comes with zero configuration! Install jump, integrate it to your shell
and let it learn your habits for a while. Simply `cd` to your directories like
you always do. After a while, jump would know how to get you when you type
`j somewhere` or just `j some`.

Maybe you made a typo like `j ssome`? No problem, jump uses fuzzy searching, so
you can type tiny, loose search term and be tolerated for your typos.

### Integration

Jump needs to be integrated into your shell to observe your `cd` habits. The
integrations also provides the `j` helper, which you would use to interact with
jump.

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
can just `j` to your projects from everywhere. 👀

![demo](https://raw.githubusercontent.com/gsamokovarov/jump/master/assets/demo.gif)

But `j` is not my favourite letter! This is fine, you can bind jump to `z`,
with this:

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

## Installation

Jump comes in packages for macOS (homebrew) and Linux.

## macOS

```bash
brew install jump
```

### Ubuntu/Debian

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.10.0/jump_0.10.0_amd64.deb
sudo dpkg -i jump_0.10.0_amd64.deb
```

### Red Hat/Fedora

```bash
wget https://github.com/gsamokovarov/jump/releases/download/v0.10.0/jump-0.10.0-1.x86_64.rpm
sudo rpm -i jump-0.10.0-1.x86_64.rpm
```

You can also build jump by yourself. Or hack on it, you know, if you like Go
and UNIX stuff. 💻

## Why

Why use jump over autojump, z or something else is a valid question. I was an
avid autojump user before building jump myself. My sloppy fingers were the main
motivation.

I mistype a lot. With autojump, I was never tolerated for a typo. I also wanted
to utilize fuzzy searching, as it saves so much effort.

Over the time I have tweaked the ranking and matching algorithm to fit my
needs, so I thing it may fit yours as well. Here is a [conversation] about
little tips and tricks using jump.

## License

Jump is licensed under the [MIT license].

✌️ <a href="https://travis-ci.org/gsamokovarov/jump">
  <img src="https://travis-ci.org/gsamokovarov/jump.svg?branch=master" alt="Build Status" data-canonical-src="https://travis-ci.org/gsamokovarov/jump.svg?branch=master">
</a>

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[conversation]: https://twitter.com/hkdobrev/status/838398833419767808
[MIT license]: https://github.com/gsamokovarov/jump/blob/master/LICENSE.txt
