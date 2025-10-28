package services

import (
	db "tmbot/database"
)

type Services struct {
	Task  taskMethods
	Party partyMethods
}

func NewServices(database db.Database) Services {
	return Services{}
}

type taskMethods interface {
	Save()
}

type partyMethods interface {
	Get()
}
