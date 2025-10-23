package postgres

import (
	"tg-bot/pkg/models"
)

func (pg *PostgresProvider) SaveTask(owner, title, desc string) error {

	_, err := pg.DB.Exec("INSERT INTO task(title, description, owner) values($1, $2, $3)", title, desc, owner)

	if err != nil {
		return err
	}

	return nil
}

func (pg *PostgresProvider) Tasks(owner string) ([]models.Task, error) {
	tasks := []models.Task{}

	err := pg.DB.Select(&tasks, "SELECT * FROM task WHERE owner=$1", owner)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
