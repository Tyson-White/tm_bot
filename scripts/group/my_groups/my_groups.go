package my_groups

import "tg-bot/types"

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {

	return &Command{ScriptInitParams: params}
}

func (cmd *Command) Run() {
	
}
