package base

import (
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
)

type StartCMD struct{ scripts.Script }

func Start(params scripts.Script) scripts.ScriptMethods {
	return &StartCMD{Script: params}
}

func (cmd *StartCMD) Run() error {
	cmd.Msg(messages.StartMSG, "./assets/welcome.png")
	cmd.Msg(messages.InfoMSG, "")

	return nil
}
