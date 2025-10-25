package create_task

import (
	"fmt"
	"strconv"
	"tg-bot/scripts"
	"tg-bot/scripts/task"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (com *Command) Run() {
	// TODO: Добавить проверку, что человек состоит в группе, в которую хочет отправить задачу

	// image := "./assets/create_task.png"

	updTitle, err := scripts.Input(com.Client, com.Session, task.TaskTitleMSG)

	if err != nil {
		return
	}

	updText, err := scripts.Input(com.Client, com.Session, task.TaskTextMSG)

	if err != nil {
		return
	}

	updGroup, err := scripts.Input(com.Client, com.Session, task.TaskGroupMSG)

	if err != nil {
		return
	}

	t, saveErr := com.Storage.SaveTask(
		com.Session.User.Username,
		updTitle.Message.Text, updText.Message.Text,
		updGroup.Message.Text,
	)

	if saveErr != nil {
		com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), task.CreateTaskErrorMSG)
		return
	} else {
		com.Client.SendMessage(strconv.Itoa(com.Session.User.ID), task.CreateTaskSuccessMSG)
	}

	usersIntoGroup, err := com.Storage.UsersByGroup(updGroup.Message.Text)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range usersIntoGroup {
		fmt.Println(id)

		// if com.Session.User.ID == id {
		// 	continue
		// }

		msg := fmt.Sprintf(`
		
		Новая задача для группы *%v*

		%v
		
		`, updGroup.Message.Text, t.ToString())

		com.Client.SendFMessage(strconv.Itoa(id), msg)

	}
}
