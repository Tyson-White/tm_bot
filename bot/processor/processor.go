package processor

import (
	"context"
	"tmbot/bot/fetcher"
	cl "tmbot/client"
	serv "tmbot/services"
	"tmbot/types"
)

type processor struct {
	services serv.Services
	sessions map[int64]*types.Session
	client   cl.Client
}

func New(services serv.Services, client cl.Client) processor {

	return processor{
		services: services,
		client:   client,
		sessions: make(map[int64]*types.Session),
	}
}

func (p processor) Handle(updPool fetcher.IncomeUpdatesPool) error {

	go func() {
		for {
			incomeUpd := <-updPool
		}
	}()
	return nil
}

func (p processor) useCommand(command string, update cl.Update) {

	switch command {
	case "start":
		go p.services.Task.Start(p.generateSession(update))
	default:
		p.client.SendError()
	}
}

func (p processor) generateSession(update cl.Update) *types.Session {
	context, cancel := context.WithCancel(context.Background())

	return &types.Session{
		Ctx:       context,
		CtxCancel: cancel,
		User:      update.Message.From,
		Sender:    cl.NewSender(update.Message.Chat.ID, p.client),
	}
}
