package router

import (
	"log"
	"strings"
	"tg-bot/client/telegram"
	"tg-bot/scripts/group/create_group"
	"tg-bot/scripts/start"
	"tg-bot/scripts/task/all_tasks"
	"tg-bot/scripts/task/create_task"
	"tg-bot/storage"
	"tg-bot/types"
)

var commands = map[string]types.ScriptFunc{
	"/create_task":  create_task.New,
	"/start":        start.New,
	"/create_group": create_group.New,
	"/tasks":        all_tasks.New,
}

type Router struct {
	sessions map[int]*types.Session
	storage  storage.Storage
	client   telegram.Client
}

func New(storage storage.Storage, client telegram.Client) Router {
	return Router{
		sessions: map[int]*types.Session{},
		storage:  storage,
		client:   client,
	}
}

func (r *Router) Start(ch <-chan []telegram.UpdateEntity) {
	log.Println("Router was initialized")

	for {
		updates := <-ch
		for _, upd := range updates {

			if strings.HasPrefix(upd.Message.Text, "/") {
				r.commandsHandler(upd)
			} else {
				s, ok := r.sessions[upd.Message.From.ID]

				if ok {
					if !s.Closed {
						s.In <- upd
					}
				}
			}
		}
	}

}

func (r *Router) commandsHandler(update telegram.UpdateEntity) {
	s, sessionExists := r.sessions[update.Message.From.ID]

	if sessionExists {
		s.CancelCtx()
	}

	script, scriptExists := commands[update.Message.Text]

	if scriptExists {
		// создаем новую сессию
		session := NewSession(update.Message.From)

		// записываем для пользователя новую сессию
		r.sessions[update.Message.From.ID] = session

		handler := script(types.ScriptInitParams{
			Session: session,
			Storage: r.storage,
			Client:  r.client,
		})

		go func() {
			defer func() {
				session.CancelCtx()
				delete(r.sessions, update.Message.From.ID)
			}()
			handler.Run()
		}()
	}
}
