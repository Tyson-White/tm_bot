package client

import (
	tg "tmbot/client/telegram"
)

type Client struct {
	Telegram methods
}

func NewClient() Client {
	return Client{
		Telegram: tg.NewTelegramClient(),
	}
}

type methods interface {
	Register(apiBotToken string) error
	Updates()
	SendMessage()
	SendPhoto()
}
