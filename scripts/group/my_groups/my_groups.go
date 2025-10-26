package my_groups

import (
	"strconv"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {

	return &Command{ScriptInitParams: params}
}

func (cmd *Command) Run() {
	groups, _ := cmd.Storage.MyGroups(cmd.Session.User.Username)

	msg := "<b>Ты состоишь в следующих группах:</b>"

	cmd.Client.SendPhoto(strconv.Itoa(cmd.Session.User.ID), "./assets/groups.png", msg)

	for _, grp := range groups {
		cmd.Client.SendFMessage(strconv.Itoa(cmd.Session.User.ID), grp.ToString())
	}
}
