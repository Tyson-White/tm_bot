package services

import (
	db "tmbot/database"
	"tmbot/services/party"
	"tmbot/services/task"
)

type Services struct {
	Task  *task.TaskService
	Party *party.Party
}

func NewServices(database db.Database) Services {
	return Services{
		Task:  task.New(),
		Party: party.New(),
	}
}
