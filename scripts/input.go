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

type InputParams struct {
	Client    telegram.Client
	Session   *types.Session
	Msg       string
	PhotoPath string
}

// TODO: Обработать закрытие контекста
func Input(params InputParams) (telegram.UpdateEntity, error) {
	ctx, cancel := context.WithCancel(context.Background())

	if params.PhotoPath == "" {
		params.Client.SendFMessage(strconv.Itoa(params.Session.User.ID), params.Msg)
	} else {
		params.Client.SendPhoto(strconv.Itoa(params.Session.User.ID), params.PhotoPath, params.Msg)
	}

	go func() {
		time.Sleep(60 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		params.Client.SendFMessage(strconv.Itoa(params.Session.User.ID), ExpiredSessionMSG)
		return telegram.UpdateEntity{}, ErrSessionTimeout
	case upd, opened := <-params.Session.In:
		if opened {

			return upd, nil
		}
	}

	return telegram.UpdateEntity{}, ErrSessionClosed

}
