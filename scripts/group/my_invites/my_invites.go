package my_invites

import (
	"log"
	"strconv"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (com *Command) Run() {
	invites, err := com.Storage.MyInvites(com.Session.User.Username)

	if err != nil {
		log.Println(err)
		com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), ErrInvites)
		return
	}

	msg := "Ваши пришлашения"

	for _, inv := range invites {
		msg += inv.ToString()
	}

	com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), msg)
}
