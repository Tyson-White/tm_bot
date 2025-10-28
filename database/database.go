package database

import "tmbot/database/postgres"

type Database struct {
	dbMethods
}

func NewDatabase() Database {
	return Database{
		dbMethods: postgres.NewPostgresDB(),
	}
}

type dbMethods interface {
	Connect()
}
