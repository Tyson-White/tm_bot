package bot

import (
	"tmbot/bot/fetcher"
	"tmbot/bot/processor"
	cl "tmbot/client"
	sv "tmbot/services"
)

type Bot struct {
	client   cl.Client
	services sv.Services
}

func NewBot(client cl.Client, services sv.Services) Bot {
	return Bot{
		client:   client,
		services: services,
	}
}

func (b Bot) Start(token string) error {

	err := b.client.Telegram.Register(token)

	if err != nil {
		return err
	}

	return b.listen()
}

func (b Bot) listen() error {

	f := fetcher.NewFetcher(b.client)
	ch := f.Fetch()

	processor.NewProcessor(b.services, ch)

	return nil
}
