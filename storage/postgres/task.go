package postgres

import (
	"tg-bot/pkg/models"

	"github.com/jmoiron/sqlx"
)

func (pg *PostgresProvider) SaveTask(owner, title, desc, group string) (models.Task, error) {
	var row *sqlx.Row

	if group != "-" {
		row = pg.DB.QueryRowx(`
			INSERT INTO task(title, description, owner) values($1, $2, $3)
			RETURNING id, title, description, created_at, groupname, owner
		`, title, desc, owner)
	} else {
		row = pg.DB.QueryRowx(`
			INSERT INTO task(title, description, owner, groupname) values($1, $2, $3, $4)
			RETURNING id, title, description, created_at, groupname, owner
		`, title, desc, owner, group)
	}

	var task models.Task

	err := row.StructScan(&task)

	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (pg *PostgresProvider) Tasks(owner string) ([]models.Task, error) {
	tasks := []models.Task{}

	err := pg.DB.Select(&tasks, "SELECT * FROM task WHERE owner=$1", owner)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
