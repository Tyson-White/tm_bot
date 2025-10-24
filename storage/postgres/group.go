package postgres

import (
	"log"
	"tg-bot/pkg/models"
)

func (p *PostgresProvider) CreateGroup(name, creator string) (models.TaskGroup, error) {
	row := p.DB.QueryRowx(`
		INSERT INTO task_group(name, creator) VALUES($1, $2)
		RETURNING id, name, creator
	`, name, creator)

	var data models.TaskGroup

	err := row.StructScan(&data)

	if err != nil {
		return models.TaskGroup{}, err
	}

	p.AddUserToGroup(creator, data.Name)

	return data, nil
}

func (p *PostgresProvider) Groups(username string) ([]models.TaskGroup, error) {
	var data []models.TaskGroup

	err := p.DB.Select(&data, "SELECT * FROM task_group WHERE creator=$1", username)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (p *PostgresProvider) AddUserToGroup(username string, group string) (bool, error) {
	// TODO: Добавить проверку, что пользователь приглашен в эту группу
	_, err := p.DB.Exec(`
		INSERT INTO users_group(username, groupname) VALUES($1, $2)
	`, username, group)

	if err != nil {
		return false, err
	}

	return true, nil

}

func (p *PostgresProvider) UsersByGroup(group string) ([]int, error) {
	var data []string

	err := p.DB.Select(&data, "SELECT username FROM users_group WHERE groupname=$1", group)

	if err != nil {
		return nil, err
	}

	var ids = []int{}

	for _, usn := range data {
		var id int

		err := p.DB.Get(&id, "SELECT telegram_id FROM t_user WHERE username=$1", usn)

		if err != nil {
			log.Println(err, usn)
			continue
		}

		ids = append(ids, id)
	}

	return ids, nil
}
