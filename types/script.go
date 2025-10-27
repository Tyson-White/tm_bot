package types

import (
	"context"
	"tg-bot/client/telegram"
)

type Session struct {
	ID        string
	User      telegram.UserEntity
	In        chan telegram.UpdateEntity
	Ctx       context.Context
	CancelCtx context.CancelFunc
	Closed    bool
}
