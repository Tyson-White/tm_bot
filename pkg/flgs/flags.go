package flgs

import (
	"flag"
	"log"
)

func MustToken() string {

	token := flag.String("bot", "", "telegram bot api token")

	flag.Parse()

	if *token == "" {
		log.Fatalln("Bot token is empty")
	}

	return *token
}
