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
	params := scripts.InputParams{
		Client:    com.Client,
		Session:   com.Session,
		Msg:       task.TaskTitleMSG,
		PhotoPath: "./assets/create_task.png",
	}
	// TODO: Добавить проверку, что человек состоит в группе, в которую хочет отправить задачу

	updTitle, err := scripts.Input(params)
	if err != nil {
		return
	}

	params.Msg = task.TaskTextMSG
	updText, err := scripts.Input(params)
	if err != nil {
		return
	}

	params.Msg = task.TaskGroupMSG
	updGroup, err := scripts.Input(params)
	if err != nil {
		return
	}

	t, saveErr := com.Storage.SaveTask(
		com.Session.User.Username,
		updTitle.Message.Text, updText.Message.Text,
		updGroup.Message.Text,
	)

	if saveErr != nil {
		com.Client.SendFMessage(strconv.Itoa(com.Session.User.ID), task.CreateTaskErrorMSG)
		return
	} else {
		com.Client.SendPhoto(strconv.Itoa(com.Session.User.ID), "./assets/success.png", task.CreateTaskSuccessMSG)
		com.Client.SendFMessage(strconv.Itoa(com.Session.User.ID), t.ToString())
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
		
		Новая задача для группы <b>%v</b>
		%v
		
		`, updGroup.Message.Text, t.ToString())

		com.Client.SendFMessage(strconv.Itoa(id), msg)

	}
}
