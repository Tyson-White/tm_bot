package postgres

import "tg-bot/pkg/models"

type Invite struct {
	Id int `db:"id"`
}

func (p *PostgresProvider) CreateInvite(groupname, creator, invited string) (int, error) {
	// Группа существует
	// Creator это вдажелец группы

	var group models.TaskGroup

	err := p.DB.Get(&group, `SELECT * FROM task_group WHERE name=$1 AND creator=$2`, groupname, creator)

	if err != nil {
		return 0, err
	}

	row := p.DB.QueryRowx("INSERT INTO invite(groupname, creator, invited) VALUES($1, $2, $3) RETURNING id", groupname, creator, invited)

	var invite Invite

	if err := row.StructScan(&invite); err != nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	return invite.Id, nil
}

func (p *PostgresProvider) InvitesByCreator(username string) ([]models.Invite, error) {
	var data []models.Invite

	err := p.DB.Select(&data, "SELECT * FROM invite WHERE creator=$1", username)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PostgresProvider) MyInvites(username string) ([]models.Invite, error) {
	var data []models.Invite

	err := p.DB.Select(&data, "SELECT * FROM invite WHERE invited=$1", username)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PostgresProvider) InviteById(id int, username string) (models.Invite, error) {
	var data models.Invite

	err := p.DB.Get(&data, "SELECT * FROM invite WHERE id=$1 AND invited=$2", id, username)

	if err != nil {
		return models.Invite{}, err
	}

	return data, nil
}
