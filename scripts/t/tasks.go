package t

import (
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
)

type UserTasksCMD struct{ scripts.Script }

func UserTasks(params scripts.Script) scripts.ScriptMethods {
	return &UserTasksCMD{Script: params}
}
func (cmd *UserTasksCMD) Run() error {

	groupUpdate, err := cmd.RequestInput(messages.ReqTaskGroup, "")

	if err != nil {
		return err
	}

	tasks, err := cmd.Storage.Tasks(cmd.Session.User.Username, groupUpdate.Message.Text)

	if err != nil {
		return err
	}

	msg := "<b>📋 Вот твои задачи: </b>\n"
	cmd.Msg(msg, "./assets/tasks.png")

	for _, t := range tasks {
		cmd.Msg(t.ToString(), "")

	}

	return nil

}
