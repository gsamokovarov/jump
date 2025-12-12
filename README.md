<p align="right">
  <a href="https://github.com/gsamokovarov/jump/releases">[releases]</a>
</p>

# Jump

**Jump** integrates with your shell and learns where you go. It tracks the directories you visit and lets you jump to the right one with just a few fuzzy-typed characters.

![Demo](./assets/demo.svg)

---

# Installation

Packages are available on the following platforms:

| Platform | Command |
| --- | --- |
| macOS | `brew install jump` or `port install jump` |
| Linux | `sudo snap install jump` |
| Nix | `nix-env -iA nixpkgs.jump` |
| Go | `go install github.com/gsamokovarov/jump@latest` |

<details>
<summary>Linux distribution packages</summary>

| Distribution | Command |
| --- | --- |
| Arch | `sudo yay -S jump` |
| Ubuntu / Debian | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump_0.51.0_amd64.deb && sudo dpkg -i jump_0.51.0_amd64.deb` |
| Fedora | `wget https://github.com/gsamokovarov/jump/releases/download/v0.51.0/jump-0.51.0-1.x86_64.rpm && sudo rpm -i jump-0.51.0-1.x86_64.rpm` |
| Void | `xbps-install -S jump` |

</details>

---

# Shell Integration

Jump is used through its helper function – `j` by default. Add it to your shell:

**bash / zsh**

```bash
eval "$(jump shell)"
```

**fish**

```
jump shell fish | source
```

**PowerShell**

```
Invoke-Expression (&jump shell pwsh | Out-String)
```

**Nushell**

```
jump shell nushell | save --append $nu.config-path
```

**Murex**

```
murex-package install https://github.com/lmorg/murex-module-jump.git
```

Jump begins tracking directories automatically once integrated.

## Custom binding

The letter `j` is the default binding for Jump. But if that's not your jam, you can customize it to your liking.

```bash
eval "$(jump shell --bind=z)"
eval "$(jump shell --bind=goto)"
```

Or bind it directly to `cd`:

```bash
eval "$(jump shell --bind=cd)"
```

Typing `cd project` now performs a fuzzy jump.

---

# Usage

## Summary (Quick Examples)

```bash
j wc               # Fuzzy Jump
j web-console      # Exact Match
j dev/soc/web      # Deep Jump
j ch site          # Multi-Part Match
j app/             # Relative Jump
j .                # Repo Root
j . cable          # Based Mode
j -                # cd -
j ..               # cd ..
j ../..            # cd ../..
j                  # Reverse Jump
j Dev              # Case-Sensitive
```

That was a quick overview of how `j` behaves. The sections below explain each feature with landing paths.

---

## Regular jump

```bash
j wc               # -> /Users/genadi/Development/rails/web-console
j console          # -> /Users/genadi/Development/rails/web-console
j b-c              # -> /Users/genadi/Development/rails/web-console
j web-console      # -> /Users/genadi/Development/rails/web-console   (exact)
```

---

## Deep jump

```bash
j ch site          # -> /Users/genadi/Development/chaos/website
j dev/soc/web      # -> /Users/genadi/Development/society/website
```

Spaces and slashes work interchangeably.

---

## Based Mode

```bash
j . cable          # -> /Users/genadi/Development/rails/rails/actioncable
j . actionview/app # -> /Users/genadi/Development/rails/rails/actionview/app
j .                # -> /Users/genadi/Development/rails/rails
```

Useful for monorepos with many repeated directory names.

---

## Relative Jump

If your input contains or ends with a slash, Jump checks for a relative directory.

```bash
# In /Users/genadi/Development/rails/rails
j actioncable/     # -> ./actioncable
j actionpack/app   # -> ./actionpack/app
```

Jump also mirrors familiar shell movements:

```bash
j -                # -> previous directory
j ..               # -> parent directory
j ../..            # -> grandparent directory
```

If the relative path does not exist, Jump falls back to fuzzy search.

---

## Reverse jump

```bash
j web              # -> /Users/genadi/Development/society/website
j                  # -> /Users/genadi/Development/chaos/website
```

---

## Case-sensitive jump

```bash
j Dev              # -> /Users/genadi/Development
```

---

## Pins

```bash
cd /Users/genadi/development/rails
jump pin r

j r                # -> /Users/genadi/development/rails
```

List and remove pins:

```bash
jump pins
jump unpin r
```

---

## Database Tools

```bash
jump clean         # remove stale entries
jump forget        # forget current directory
jump top           # ranked directories
jump top --score   # ranked directories with numeric scores
jump top dev       # fuzzy filtered list
```

---

## Importing from `autojump` or `z`

```bash
jump import
jump import autojump
jump import z
```

Imports merge into your existing database without overwriting scores.

---

# Why Jump?

Jump embraces fuzzy matching, forgiving inputs, expressive hints (`Dev`, `soc/web`, tapping `j` again), and natural relative path handling.

It is designed to save keystrokes without giving up control.

If it feels hand-crafted… that is for you to call. 😄
