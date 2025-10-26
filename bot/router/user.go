package router

import (
	"log"
	"tg-bot/client/telegram"
)

func (r *Router) defineUser(user telegram.UserEntity) error {
	userExists, err := r.storage.UserExists(user.ID)

	if err != nil {
		log.Println(err)
	}

	if !userExists {
		err := r.storage.SaveUser(user.ID, user.Username)

		if err != nil {
			log.Printf("New user in system: %v", user.Username)
		}
	}

	return nil
}
