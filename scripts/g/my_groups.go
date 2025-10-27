package g

import (
	"tg-bot/scripts"
)

type UserGroupsCMD struct{ scripts.Script }

func UserGroups(params scripts.Script) scripts.ScriptMethods {
	return &UserGroupsCMD{Script: params}
}
func (cmd *UserGroupsCMD) Run() error {
	groups, _ := cmd.Storage.MyGroups(cmd.Session.User.Username)

	msg := "<b>Ты состоишь в следующих группах:</b>"

	cmd.Msg(msg, "./assets/groups.png")

	for _, grp := range groups {
		cmd.Msg(grp.ToString(), "")
	}

	return nil
}
