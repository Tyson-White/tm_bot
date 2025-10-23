package scripts

import (
	"context"
	"errors"
	"strconv"
	"tg-bot/client/telegram"
	"time"
)

var ErrSessionClosed = errors.New("session is closed")
var ErrSessionTimeout = errors.New("session is closed")

// TODO: Обработать закрытие контекста
func Input(client telegram.Client, chatId int, ch chan telegram.UpdateEntity, msg string) (telegram.UpdateEntity, error) {
	ctx, cancel := context.WithCancel(context.Background())

	client.SendMessage(strconv.Itoa(chatId), msg)

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		client.SendMessage(strconv.Itoa(chatId), ExpiredSessionMSG)
		return telegram.UpdateEntity{}, ErrSessionTimeout
	case upd, opened := <-ch:
		if opened {

			return upd, nil
		}
	}

	return telegram.UpdateEntity{}, ErrSessionClosed

}
