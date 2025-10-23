package main

import (
	"tg-bot/bot"
	"tg-bot/client/telegram"
	"tg-bot/pkg/flgs"
	"tg-bot/storage"
	"tg-bot/storage/postgres"
)

func main() {

	pg := postgres.New()
	storage := storage.New(&pg)

	tgClient := telegram.New(flgs.MustToken())

	b := bot.New(tgClient, storage, 2000)
	b.Run()
}
