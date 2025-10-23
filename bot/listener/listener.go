package listener

import (
	"log"
	"tg-bot/client/telegram"
	"time"
)

type Listener struct {
	lastUpdateID int
	timeSleep    int
	client       telegram.Client
}

func New(timeSleep int, client telegram.Client) Listener {
	return Listener{
		timeSleep: timeSleep,
		client:    client,
	}
}

// Получает апдейты каждые сколько-то секунд
func (l *Listener) Listen(ch chan<- []telegram.UpdateEntity) {
	log.Println("Listner was initialized")

	for {
		func() {
			defer time.Sleep(time.Duration(l.timeSleep) * time.Millisecond)

			upds, err := l.client.Updates(l.lastUpdateID)

			if err != nil {
				log.Println("Telegram updates error:", err)
				return
			}

			if len(upds) > 0 {
				// Ставим в запрос offset начиная с последнего полученного апдейта + 1
				l.lastUpdateID = upds[len(upds)-1].ID + 1
			}

			ch <- l.filterUpdates(upds)

		}()
	}
}

func (l *Listener) filterUpdates(updates []telegram.UpdateEntity) []telegram.UpdateEntity {
	filtred := []telegram.UpdateEntity{}

	for _, upd := range updates {
		updTime := time.Unix(int64(upd.Message.Date), 0)
		updExpireTime := updTime.Add(time.Duration(l.timeSleep) * time.Millisecond)

		if updExpireTime.After(time.Now()) {
			filtred = append(filtred, upd)
		}
	}

	return filtred
}
