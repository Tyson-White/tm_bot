package processor

import (
	"tmbot/bot/fetcher"
	serv "tmbot/services"
)

type processor struct {
	services serv.Services
}

func NewProcessor(services serv.Services, in fetcher.Channel) processor {

	return processor{
		services: services,
	}
}

func (p processor) Handle(input chan []string) error {
	return nil
}
