package storage

import (
	"tg-bot/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB *sqlx.DB
	StorageMethods
}

func New(provider StorageMethods) Storage {
	return Storage{
		DB:             provider.Connect(),
		StorageMethods: provider,
	}
}

type StorageMethods interface {
	Connect() *sqlx.DB
	SaveTask(owner, title, desc string) error
	Tasks(user string) ([]models.Task, error)
	// DeleteTask()
	// CompleteTask()

	// CreateGroup()
	// DeleteGroup()

	// CreateInvite()
	// AcceptInvite()
}
