package postgres

import (
	"errors"
	"log"
	"strings"
	"tg-bot/pkg/e"
	"tg-bot/pkg/models"

	"github.com/jmoiron/sqlx"
)

func (pg *PostgresProvider) SaveTask(owner, title, desc, group string) (models.Task, error) {
	var row *sqlx.Row

	if group == "skip" {
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
		if strings.Contains(err.Error(), "constraint") {
			return models.Task{}, e.ErrGroupNotFound
		}

		return models.Task{}, errors.Join(e.ErrServerError, err)
	}

	return task, nil
}

func (pg *PostgresProvider) Tasks(user string, groupname string) ([]models.Task, error) {
	tasks := []models.Task{}

	err := pg.DB.Select(&tasks, `
	SELECT DISTINCT t.* FROM task AS t 
	JOIN users_group AS ug ON t.groupname=ug.groupname WHERE ug.groupname = $1 AND ug.username = $2
             `, groupname, user)

	if err != nil {
		return nil, errors.Join(e.ErrServerError, err)
	}

	log.Println(len(tasks))

	return tasks, nil
}
