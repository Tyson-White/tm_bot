package scripts

import (
	"context"
	"errors"
	"strconv"
	"tg-bot/client/telegram"
	"tg-bot/pkg/e"
	"tg-bot/pkg/messages"
	"tg-bot/storage"
	"tg-bot/types"
	"time"
)

type ScriptFunc = func(script Script) ScriptMethods

type Script struct {
	Client  telegram.Client
	Session *types.Session
	Storage storage.Storage
}

func New(client telegram.Client, session *types.Session) Script {
	return Script{
		Client:  client,
		Session: session,
	}
}

type ScriptMethods interface {
	Run() error
}

func (sc *Script) RequestInput(msg string, filepath string) (telegram.UpdateEntity, error) {
	ctx, cancel := context.WithCancel(context.Background())

	sc.Msg(msg, filepath)

	go func() {
		time.Sleep(60 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		return telegram.UpdateEntity{}, errors.Join(e.ErrClientError, errors.New(messages.ExpiredSession))
	case upd, opened := <-sc.Session.In:
		if opened {

			return upd, nil
		}
	}

	return telegram.UpdateEntity{}, e.ErrSessionClosed

}

func (sc *Script) Msg(msg, filepath string) (telegram.MessageEntity, error) {
	strId := strconv.Itoa(sc.Session.User.ID)

	if filepath == "" {
		return sc.Client.SendFMessage(strId, msg)
	} else {
		return sc.Client.SendPhoto(strId, filepath, msg)
	}
}

func (sc *Script) Err(msg string) (telegram.MessageEntity, error) {
	strId := strconv.Itoa(sc.Session.User.ID)

	return sc.Client.SendError(strId, msg)
}

func (sc *Script) Success(msg string) (telegram.MessageEntity, error) {
	strId := strconv.Itoa(sc.Session.User.ID)

	return sc.Client.SendSuccess(strId, msg)
}
