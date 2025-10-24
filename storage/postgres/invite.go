package postgres

import "tg-bot/pkg/models"

func (p *PostgresProvider) CreateInvite(groupId int, creator, invited string) (bool, error) {

	_, err := p.DB.Exec(`
		INSERT INTO invite(group_id, creator, invited) VALUES($1, $2, $3)
	`, groupId, creator, invited)

	if err != nil {
		return false, err
	}

	return true, nil
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
