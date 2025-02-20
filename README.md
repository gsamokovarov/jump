<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

**Jump** integrates with your shell and learns about your navigational habits by
keeping track of the directories you visit. It gives you the most visited
directory for the shortest search term you type.

![Demo](./assets/demo.svg)

## Installation

Jump comes in packages for the following platforms.

| Platform | Command |
| --- | --- |
| macOS | `brew install jump` or `port install jump` |
| Linux | `sudo snap install jump` |
| Nix | `nix-env -iA nixpkgs.jump` |
| Go | `go install github.com/gsamokovarov/jump@latest` |

<details>
<summary>Linux distribution specific packages</summary>

| Distribution | Command |
| --- | --- |
| Void | `xbps-install -S jump` |
| Ubuntu | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump_0.51.0_amd64.deb && sudo dpkg -i jump_0.51.0_amd64.deb` |
| Debian | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump_0.51.0_amd64.deb && sudo dpkg -i jump_0.51.0_amd64.deb` |
| Fedora | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump-0.51.0-1.x86_64.rpm && sudo rpm -i jump-0.51.0-1.x86_64.rpm` |

</details>

### Integration

You are using Jump through its shell helper function, `j`. To get it, you have
to integrate Jump with your shell.

#### bash / zsh

Add the line below in `~/.bashrc`, `~/bash_profile` or `~/.zshrc`:

```bash
eval "$(jump shell)"
```

#### fish

Add the line below in `~/.config/fish/config.fish`:

```
jump shell fish | source
```

#### PowerShell

Add the line below needs to your profile, located by typing `$PROFILE`:

```
Invoke-Expression (&jump shell pwsh | Out-String)
```

Once integrated, Jump will automatically monitor directory changes and start
building an internal database.

#### Murex

Jump bindings can be installed directly from Murex:

```
murex-package install https://github.com/lmorg/murex-module-jump.git
```

Please note that this doesn't install `jump` itself. You will still need to
install the `jump` executable using the installation instructions above.

### But `j` is not my favorite letter!

This is fine, you can bind jump to `z` with the following integration command:

```bash
eval "$(jump shell --bind=z)"
```

Typing `z dir` would just work! This is only an example, you can bind it to
_anything_. If you are one of those persons that likes to type a lot with their
fingers, you can do:

```bash
eval "$(jump shell --bind=goto)"
```

Voila! `goto dir` becomes a thing. The possibilities are endless!

## Usage

Once integrated, **jump** introduces the **j** helper. It accepts only search
terms, and as a design goal, there are no arguments for **j**. Whatever you give
it, it's treated as a search term.

**Jump** uses fuzzy matching to find the desired directory to jump to. This
means that your search terms are patterns that match the desired directory
approximately rather than exactly. Typing **2** to **5** consecutive characters
of the directory name is all that **jump** needs to find it.

### Regular jump

The default search behavior of **jump** is to match the
directory name of a score. The match is case insensitive.

If you visit the directory `/Users/genadi/Development/rails/web-console` often,
you can jump to it by:

```bash
$ j wc      # or...
$ j webc    # or...
$ j console # or...
$ j b-c     # or...
```

Using jump is all about saving keystrokes. However, if you made the effort to
type a directory base name exactly, **jump** will try to find the exact match,
rather than a fuzzy search.

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
other absolute path above. The jump above will land on the most-scored path,
which is the `society` one, however, what if we wanted to land on the `chaos`
website?

```bash
$ j ch site
$ pwd
/Users/genadi/Development/chaos/website
```

This instructs **jump** to look for a `site` match inside that is preceded by a
`ch` match in the parent directory. The search is normalized only on the last
two parts of the target paths. This will ensure a better match because of the
shorter path to a fuzzy match.

There are no depth limitations, though and a jump to
`/Users/genadi/Development/society/website` can look like this:

```bash
$ j dev soc web
$ pwd
/Users/genadi/Development/society/website
```

Every space passed to `j` is converted to an OS separator. The last
search term can be expressed as:

```bash
$ j dev/soc/web
$ pwd
/Users/genadi/Development/society/website
```

## Reverse jump

Bad jumps happen. Sometimes, we're looking for a directory that doesn't have the
best score at the moment. Let's work with the following following jump database:

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
without any arguments, will instruct **jump** to go to the second-best match.

```bash
$ j
$ pwd
/Users/genadi/Development/chaos/website
```

### Case-sensitive jump

To trigger a case-sensitive search, use a term that has a capital letter.

```bash
$ j Dev
$ pwd
/Users/genadi/Development
```

The jump will resolve to `/Users/genadi/Development` even if there is
`/Users/genadi/Development/dev-tools` that scores better.

### Pins

For various reasons, Jump may not always find the directory you want, but don't worry—you can make it work!

A pin forces an input to always go to a specific location. If you want j r to always go to /Users/genadi/development/rails, you can do:

```
$ cd /Users/genadi/development/rails
$ jump pin r
$ cd
$ j r # Skips the scoring and goes straight to the pinned directory.
$ pwd
/Users/genadi/development/rails
```

Notice the `jump` command instead of the `j` shell function helper. `j` will always treat its input as search terms. It may apply some heuristics to how the input looks, but it will never accept arguments or switches. Here is where the `jump` command comes in. It is bundled with lots of helpers to make your `j` life easier. The pins are one of them.

Try `jump --help` for all those hidden (**not** not-documented ones) features.

## Is it like autojump or z?

I was an avid autojump user, but it wasn't forgiving my sloppy fingers. That
pushed me to create Jump with the goal of accepting fuzzy search terms. This
lets you type a couple of letters and go to your project:

`j web` vs `j website`

The fuzzy typing is your fingers-friendly. You can make a typo, and the jump
would mostly work:

`j wwebsite`

As a design goal, the `j` helper doesn't have any arguments. It's all about the search
term. That said, you can use the search term itself to hint jump about the desired directory.

Typing mixed case input would force a case-sensitive match:

`j Dev` would prefer /Users/genadi/Development

If you have multiple projects with the same name in umbrella directories you
can hint with OS separators:

`j soc/web` -> /society/website
`j ra/web` -> /raketa/website

If your input doesn’t give you the right dir, you can `j`. That will jump to
the next entry with the previous input.

Little hand-tuned details like those let `jump` read my mind with zero LLMs
interaction. If I wasn't a humble developer, I'd call it artisan
hard-crafted software, but I am, so I'll let you call it what you want. 😄

## Migrate from `autojump` or `z`

You can import your datafile from `autojump` or `z` with:

```bash
$ jump import
```

This will try `z` first, then `autojump`, so you can even combine all the
entries from both tools.

The command is safe to run on a pre-existing jump database, because if an entry
exist in jump already, it won't be imported, and its score will remain
unchanged. You can be explicit and choose to import `autojump` or `z` with:

```bash
$ jump import autojump
$ jump import z
```

## Thanks! 🙌

Thank you for stopping by and showing your interest in Jump!

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[conversation]: https://twitter.com/hkdobrev/status/838398833419767808
