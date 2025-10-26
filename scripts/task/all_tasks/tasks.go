package all_tasks

import (
	"fmt"
	"strconv"
	"tg-bot/scripts/task"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (c *Command) Run() {

	tasks, err := c.Storage.Tasks(c.Session.User.Username)

	if err != nil {
		fmt.Println(err)
		c.Client.SendFMessage(strconv.Itoa(c.Session.User.ID), task.GetTasksErrorMSG)
	}

	msg := "<b>📋 Вот твои задачи: </b>\n"
	c.Client.SendPhoto(strconv.Itoa(c.Session.User.ID), "./assets/tasks.png", msg)

	for _, t := range tasks {
		c.Client.SendFMessage(strconv.Itoa(c.Session.User.ID), t.ToString())

	}

}
