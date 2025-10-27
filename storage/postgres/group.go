package postgres

import (
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

func (p *PostgresProvider) MyGroups(username string) ([]models.TaskGroup, error) {
	var groups []models.TaskGroup

	err := p.DB.Select(&groups, `
	SELECT tg.id, tg.name, tg.creator FROM task_group AS tg
	JOIN users_group ON tg.name=users_group.groupname 
	WHERE users_group.username=$1
	`, username)

	if err != nil {
		return nil, err
	}

	return groups, nil

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
	var data []int

	err := p.DB.Select(&data, `
	SELECT t_user.telegram_id FROM t_user 
    JOIN users_group ON t_user.username=users_group.username WHERE users_group.groupname=$1
    `, group)

	if err != nil {
		return nil, err
	}

	return data, nil
}
