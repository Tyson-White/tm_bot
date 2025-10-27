package inv

import (
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
)

type UserInvitesCMD struct{ scripts.Script }

func UserInvites(params scripts.Script) scripts.ScriptMethods {
	return &UserInvitesCMD{Script: params}
}
func (cmd *UserInvitesCMD) Run() error {
	invites, err := cmd.Storage.MyInvites(cmd.Session.User.Username)

	if err != nil {
		cmd.Err(messages.ErrInvites)
		return err
	}

	msg := "Ваши пришлашения"

	for _, inv := range invites {
		msg += inv.ToString()
	}

	cmd.Msg(msg, "")

	return nil
}
