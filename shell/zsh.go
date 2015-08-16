package shell

var Zsh = Shell(`# A chpwd hook for zsh.
__jump_chpwd() {
  jump update
}

# Now, add our function to the chpwd_functions list.
typeset -gaU chpwd_functions
chpwd_functions+=__jump_chpwd

# Shortcut jump to j for the autojump folks.
j() {
  local dir="$(jump cd $@)"
  test -d "$dir" && cd "$dir"
}
`)
