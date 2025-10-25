package router

import (
	"log"
	"strings"
	"tg-bot/client/telegram"
	"tg-bot/scripts/group/accept_invite"
	"tg-bot/scripts/group/create_group"
	"tg-bot/scripts/group/my_invites"
	"tg-bot/scripts/start"
	"tg-bot/scripts/task/all_tasks"
	"tg-bot/scripts/task/create_task"
	"tg-bot/storage"
	"tg-bot/types"
)

var commands = map[string]types.ScriptFunc{
	"/start":         start.New,
	"/create_task":   create_task.New,
	"/tasks":         all_tasks.New,
	"/create_group":  create_group.New,
	"/accept_invite": accept_invite.New,
	"/invites":       my_invites.New,
	// invite

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

			// TODO: Обработать ситуацию, когда в функции позникает ошибка
			r.defineUser(upd.Message.From)

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
