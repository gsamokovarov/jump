# Hook jump on directory changes.
function __jump_add --on-variable PWD
  status --is-command-substitution; and return
  jump update
end

# Shortcut to j for the autojump folks.
function j
  set -l output (jump cd $argv)
  if test -d "$output"
	  cd $output
  end
end
