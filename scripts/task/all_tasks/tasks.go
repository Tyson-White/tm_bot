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
		c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), task.GetTasksErrorMSG)
	}

	response := "📋 Вот твои задачи: \n"

	for _, t := range tasks {
		response += t.ToString()
	}

	c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), response)

}
