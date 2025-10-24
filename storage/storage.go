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
	SaveTask(owner, title, desc, group string) (models.Task, error)
	Tasks(user string) ([]models.Task, error)
	// DeleteTask()
	// CompleteTask()

	CreateGroup(name, creator string) (models.TaskGroup, error)
	Groups(username string) ([]models.TaskGroup, error)
	AddUserToGroup(username string, group string) (bool, error)
	UsersByGroup(group string) ([]int, error)
	// DeleteGroup()

	UserByName(username string) (models.TUser, error)

	CreateInvite(groupId int, creator, invited string) (bool, error)
	// InvitesByName(username string) ([]models.Invite, error)
	InviteById(id int, username string) (models.Invite, error)
	MyInvites(username string) ([]models.Invite, error)

	SaveUser(id int, username string) error
	UserExists(id int) (bool, error)
	// AcceptInvite()
}
