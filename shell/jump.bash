# A PROMPT_COMMAND hook for bash.
__jump_prompt_command() {
    jump update "$(pwd)"
}

# Called before drawing the prompt. We can use it to hook up on directory
# changes instead of overriding the cd function.
export PROMPT_COMMAND="__jump_prompt_command"

# Shortcut jump to j for the autojump folks.
j() {
    local output="$(jump cd ${@})"
    [ -d "${output}" ] && cd "${output}"
}
