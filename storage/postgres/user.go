package postgres

import (
	"database/sql"
	"errors"
	"tg-bot/pkg/models"
)

func (p *PostgresProvider) SaveUser(id int, username string) error {

	_, err := p.DB.Exec("INSERT INTO t_user(telegram_id, username) VALUES($1, $2)", id, username)

	if err != nil {
		return err
	}

	return err
}

// false - если нет в БД, true - если есть
func (p *PostgresProvider) UserExists(id int) (bool, error) {

	var user models.TUser

	err := p.DB.Get(&user, "SELECT * FROM t_user WHERE telegram_id=$1", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *PostgresProvider) UserByName(username string) (models.TUser, error) {

	var user models.TUser

	err := p.DB.Get(&user, "SELECT * FROM t_user WHERE username=$1", username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TUser{}, nil
		}

		return models.TUser{}, err
	}

	return user, nil
}
