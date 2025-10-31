package types

import (
	"context"
	"tmbot/client"
)

type Session struct {
	Ctx        context.Context
	CtxCancel  context.CancelFunc
	IncomePool <-chan []client.Update
	User       client.User
	Sender     client.Sender
}
