package bot

import (
	"tg-bot/bot/listener"
	"tg-bot/bot/router"
	"tg-bot/client/telegram"
	"tg-bot/storage"
)

type Bot struct {
	// Сервис с которым взаимодействует бот
	Client telegram.Client

	// БД куда все складируется
	Storage storage.Storage

	// В миллисекундах
	listenerCooldown int
}

func New(
	tgClient telegram.Client,
	storage storage.Storage,
	sleepTime int,
	// discordClient
	// storage
) Bot {

	bot := Bot{
		Client:           tgClient,
		listenerCooldown: sleepTime,
		Storage:          storage,
	}

	return bot
}

func (b *Bot) Run() {
	updatesChannel := make(chan []telegram.UpdateEntity)

	l := listener.New(b.listenerCooldown, b.Client)
	go l.Listen(updatesChannel)

	r := router.New(b.Storage, b.Client)
	r.Start(updatesChannel)
}
