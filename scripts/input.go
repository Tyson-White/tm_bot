package scripts

import (
	"context"
	"errors"
	"strconv"
	"tg-bot/client/telegram"
	"tg-bot/types"
	"time"
)

var ErrSessionClosed = errors.New("session is closed")
var ErrSessionTimeout = errors.New("session is closed")

// TODO: Обработать закрытие контекста
func Input(client telegram.Client, session *types.Session, msg string) (telegram.UpdateEntity, error) {
	ctx, cancel := context.WithCancel(context.Background())

	client.SendFMessage(strconv.Itoa(session.User.ID), msg)

	go func() {
		time.Sleep(60 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		client.SendMessage(strconv.Itoa(session.User.ID), ExpiredSessionMSG)
		return telegram.UpdateEntity{}, ErrSessionTimeout
	case upd, opened := <-session.In:
		if opened {

			return upd, nil
		}
	}

	return telegram.UpdateEntity{}, ErrSessionClosed

}
