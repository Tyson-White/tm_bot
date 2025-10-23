package create_group

import (
	"fmt"
	"tg-bot/scripts"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (com *Command) Run() {

	nameUpd, err := scripts.Input(com.Client, com.Session.User.ID, com.Session.In, GroupNameMSG)

	if err != nil {
		return
	}

	usersUpd, err := scripts.Input(com.Client, com.Session.User.ID, com.Session.In, UsersMSG)

	if err != nil {
		return
	}

	fmt.Println(nameUpd.Message.Text, usersUpd.Message.Text)
}
