package create_task

import (
	"strconv"
	"tg-bot/client/telegram"
	"tg-bot/scripts"
	"tg-bot/scripts/task"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (com *Command) Run() {

	updTitle, err := scripts.Input(com.Client, com.Session.User.ID, com.Session.In, task.TaskTitleMSG)

	if err != nil {
		return
	}

	updText, err := scripts.Input(com.Client, com.Session.User.ID, com.Session.In, task.TaskTextMSG)

	if err != nil {
		return
	}

	saveErr := com.SaveTask(updTitle.Message.Text, updText.Message.Text)

	if saveErr != nil {
		com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), task.CreateTaskErrorMSG)
	} else {
		com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), task.CreateTaskSuccessMSG)
	}

}

func (com *Command) SaveTask(title, desc string) error {
	err := com.Storage.SaveTask(telegram.FormatUsername(com.Session.User.Username), title, desc)

	return err
}
