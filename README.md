<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

**Jump** integrates with your shell and learns about your navigational habits by
keeping track of the directories you visit. It gives you the most visited
directory for the shortest search term you type.

![Demo](./assets/demo.svg)

## Installation

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
| Arch | `sudo yay -S jump` |
| Ubuntu | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump_0.51.0_amd64.deb && sudo dpkg -i jump_0.51.0_amd64.deb` |
| Debian | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump_0.51.0_amd64.deb && sudo dpkg -i jump_0.51.0_amd64.deb` |
| Fedora | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump-0.51.0-1.x86_64.rpm && sudo rpm -i jump-0.51.0-1.x86_64.rpm` |
| Void | `xbps-install -S jump` |
</details>

### Integration

Integrate Jump with your shell to use the `j` helper function.

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

Add to your profile (find it with `$PROFILE`):

```
Invoke-Expression (&jump shell pwsh | Out-String)
```

#### Nushell

Append the integration to your config:

```
jump shell nushell | save --append $nu.config-path
```

Once integrated, Jump automatically monitors directory changes.

#### Murex

Install directly from Murex (this doesn't install `jump` itself):

```
murex-package install https://github.com/lmorg/murex-module-jump.git
```

### Custom Binding

Bind jump to any name you like:

```bash
eval "$(jump shell --bind=z)"      # z dir
eval "$(jump shell --bind=goto)"   # goto dir
eval "$(jump shell --bind=cd)"     # cd dir (replaces builtin cd!)
```

You can even replace the builtin `cd` command. Jump handles `j ..` and `j -` just like the shell.

## Usage

The `j` helper accepts only search terms and has no arguments. Jump uses fuzzy matching to find directories, requiring only a few consecutive characters.

### Shell Shortcuts

Jump handles common shell navigation:

```bash
$ j ..      # Same as cd ..
$ j -       # Same as cd - (previous directory)
```

### Regular Jump

Jump matches directory names using case-insensitive fuzzy search by default. For `/Users/genadi/Development/rails/web-console`:

```bash
$ j wc      # or webc, or console, or b-c
```

Exact matches take priority over fuzzy search:

```bash
$ j web-console
$ pwd
/Users/genadi/Development/rails/web-console
```

### Deep Jump

For multiple directories with similar names:

```bash
/Users/genadi/Development/society/website
/Users/genadi/Development/chaos/website
```

Use parent directory hints to distinguish them:

```bash
$ j ch site
$ pwd
/Users/genadi/Development/chaos/website
```

Chain multiple terms (spaces become path separators):

```bash
$ j dev/soc/web          # or: j dev soc web
$ pwd
/Users/genadi/Development/society/website
```

## Based Mode

Use `j .` to search within your current git repository root (or `JUMP_BASED_PATH` env var):

```bash
$ j . cable              # From anywhere in rails/rails monorepo
$ pwd
/Users/genadi/Development/rails/rails/actioncable

$ j . actioncable/app    # Direct paths work too
$ pwd
/Users/genadi/Development/rails/rails/actioncable/app

$ j .                    # Return to repository root
$ pwd
/Users/genadi/Development/rails/rails
```

## Relative Jump

If a search ends on, or cointains a slash, like `j app/`, and the directory exists
relative to the current directory, Jump will enter it just like cd would:

```bash
$ cd /Users/genadi/Development/balkan
$ j app/                 # Searches within balkan/
$ pwd
/Users/genadi/Development/balkan/app
```

## Reverse Jump

If the first match isn't what you wanted, run `j` again without arguments to try the next match:

```bash
$ j web
$ pwd
/Users/genadi/Development/society/website
$ j         # Try next match
$ pwd
/Users/genadi/Development/chaos/website
```

### Case-Sensitive Search

Use a capital letter to trigger case-sensitive matching:

```bash
$ j Dev
$ pwd
/Users/genadi/Development   # Preferred over /Users/genadi/Development/dev-tools
```

### Pins

Pin a search term to always go to a specific directory:

```bash
$ cd /Users/genadi/development/rails
$ jump pin r
$ j r                       # Always jumps to rails, skipping fuzzy search
$ pwd
/Users/genadi/development/rails
```

List all pins:

```bash
$ jump pins
r    /Users/genadi/Development/rails
w    /Users/genadi/Development/website
```

Remove a pin:

```bash
$ jump unpin r
```

## Database Management

### Clean

Remove entries for deleted directories:

```bash
$ jump clean
```

Jump runs this automatically when jumping to non-existent directories (unless `--preserve=true`).

### Forget

Remove a directory from the database:

```bash
$ jump forget              # Current directory
$ jump forget /path/to/dir # Specific path
```

### Top

View highest-scored directories:

```bash
$ jump top                 # All directories
$ jump top --score         # Show numeric scores
$ jump top dev             # Filter by search term
```

## Is it like autojump or z?

Jump was created to be more forgiving with fuzzy search. Key differences:

- **Fuzzy matching**: `j web` finds `website` (typos work too: `j wwebsite`)
- **No arguments**: Everything is a search term
- **Smart hints**: Mixed case (`j Dev`), path separators (`j soc/web`), reverse jump (`j` with no args)
- **Shell integration**: Handles `j ..`, `j -`, and can even replace `cd`

Little hand-tuned details like these let jump read your mind with zero LLMs. If I wasn't humble, I'd call it artisan hard-crafted software, but I am, so I'll let you call it what you want. 😄

## Migrate from autojump or z

Import your existing database:

```bash
$ jump import              # Tries z first, then autojump
$ jump import autojump     # Explicit import
$ jump import z
```

Safe to run on existing Jump databases—won't overwrite scores.

## Thanks! 🙌

Thank you for stopping by and showing your interest in Jump!

[man]: http://gsamokovarov.com/jump
[Go workspace]: https://golang.org/doc/code.html#Workspaces
[conversation]: https://twitter.com/hkdobrev/status/838398833419767808
