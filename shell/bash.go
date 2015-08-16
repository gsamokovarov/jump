package shell

var Bash = Shell(`# A PROMPT_COMMAND hook for bash.
__jump_prompt_command() {
  jump update
}

# Called before drawing the prompt. We can use it to hook up on directory
# changes instead of overriding the cd function.
test "$PROMPT_COMMAND" =~ __jump_prompt_command || {
  PROMPT_COMMAND="__jump_prompt_command;$PROMPT_COMMAND"
}

# Shortcut jump to j for the autojump folks.
j() {
  local dir="$(jump cd $@)"
  test -d "$dir"  && cd "$dir"
}
`)
