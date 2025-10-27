package router

import (
	"tg-bot/client/telegram"
)

func (r *Router) defineUser(user telegram.UserEntity) error {
	userExists, err := r.storage.UserExists(user.ID)

	if err != nil {
		return err
	}

	if !userExists {
		err := r.storage.SaveUser(user.ID, user.Username)
		
		if err != nil {
			return err
		}
	}

	return nil
}
