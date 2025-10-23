package scripts

import (
	"errors"
	"strconv"
	"tg-bot/client/telegram"
	"time"
)

var ErrSessionClosed = errors.New("session is closed")
var ErrSessionTimeout = errors.New("session is closed")

// TODO: Обработать закрытие контекста
func Input(client telegram.Client, chatId int, ch chan telegram.UpdateEntity, msg string) (telegram.UpdateEntity, error) {
	timeout := false

	client.SendMessage(strconv.Itoa(chatId), msg)

	go func() {
		time.Sleep(5 * time.Second)
		timeout = true
		ch <- telegram.UpdateEntity{}
	}()

	if upd, opened := <-ch; opened {

		if timeout {
			client.SendMessage(strconv.Itoa(chatId), `
🚫 Время ожидания истекло.

➡️ Введите команду снова.`)
			return upd, ErrSessionTimeout
		}

		return upd, nil
	}

	return telegram.UpdateEntity{}, ErrSessionClosed

}
