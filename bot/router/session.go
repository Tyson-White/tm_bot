package router

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"tg-bot/client/telegram"
	"tg-bot/types"
)

func NewSession(user telegram.UserEntity) *types.Session {
	ctx, cancelCtx := context.WithCancel(context.Background())
	id := rand.Int()
	s := types.Session{
		ID:        fmt.Sprintf("%v-%d", user.Username, id),
		In:        make(chan telegram.UpdateEntity),
		Ctx:       ctx,
		CancelCtx: cancelCtx,
		User:      user,
		Closed:    false,
	}

	log.Printf("[Connected] %v", s.ID)

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.Closed = true
				close(s.In)
				log.Printf("[Disconnected] %v", s.ID)
				return
			}
		}
	}()

	return &s
}
