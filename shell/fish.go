package shell

var Fish = Shell(`# Hook jump on directory changes.
function __jump_add --on-variable PWD
  status --is-command-substitution; and return
  jump update
end

# Shortcut to j for the autojump folks.
function j
  set -l dir (jump cd $argv)
  test -d "$dir"; and cd "$dir"
end
`)
