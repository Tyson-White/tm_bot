package types

import (
	"context"
	"tg-bot/client/telegram"
	"tg-bot/storage"
)

type ScriptCommandHandler interface {
	Run()
}

type ScriptInitParams struct {
	Session *Session
	Storage storage.Storage
	Client  telegram.Client
}

type ScriptFunc = func(ScriptInitParams) ScriptCommandHandler

type Session struct {
	ID        string
	User      telegram.UserEntity
	In        chan telegram.UpdateEntity
	Ctx       context.Context
	CancelCtx context.CancelFunc
	Closed    bool
}
