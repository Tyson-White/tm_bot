package t

import (
	"fmt"
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
)

type TaskCreationCMD struct{ scripts.Script }

func TaskCreation(params scripts.Script) scripts.ScriptMethods {
	return &TaskCreationCMD{Script: params}
}

func (cmd *TaskCreationCMD) Run() error {
	photo := "./assets/create_task.png"
	// TODO: Добавить проверку, что человек состоит в группе, в которую хочет отправить задачу

	updTitle, err := cmd.RequestInput(messages.ReqTaskTitle, photo)
	if err != nil {
		return err
	}

	updText, err := cmd.RequestInput(messages.ReqTaskText, photo)
	if err != nil {
		return err
	}

	updGroup, err := cmd.RequestInput(messages.ReqTaskGroupMSG, photo)
	if err != nil {
		return err
	}

	t, saveErr := cmd.Storage.SaveTask(
		cmd.Session.User.Username,
		updTitle.Message.Text, updText.Message.Text,
		updGroup.Message.Text,
	)

	if saveErr != nil {
		return saveErr
	} else {
		cmd.Msg(messages.CreateTaskSuccess, "./assets/success.png")
		cmd.Msg(t.ToString(), "")
	}

	usersIntoGroup, err := cmd.Storage.UsersByGroup(updGroup.Message.Text)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, id := range usersIntoGroup {
		fmt.Println(id)

		if cmd.Session.User.ID == id {
			continue
		}

		msg := fmt.Sprintf(`
		
		Новая задача для группы <b>%v</b>
		%v
		
		`, updGroup.Message.Text, t.ToString())

		cmd.Msg(msg, "")

	}

	return nil
}
