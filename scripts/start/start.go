package start

import (
	"strconv"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (c *Command) Run() {
	c.Client.SendFMessage(strconv.Itoa(c.Session.User.ID), StartMsg)
}
