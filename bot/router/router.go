package router

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"tg-bot/client/telegram"
	"tg-bot/pkg/e"
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
	"tg-bot/scripts/base"
	"tg-bot/scripts/g"
	"tg-bot/scripts/inv"
	"tg-bot/scripts/t"
	"tg-bot/storage"
	"tg-bot/types"
)

type Router struct {
	sessions map[int]*types.Session
	storage  storage.Storage
	client   telegram.Client
	commands map[string]scripts.ScriptFunc
}

func New(storage storage.Storage, client telegram.Client) Router {
	return Router{
		sessions: map[int]*types.Session{},
		storage:  storage,
		client:   client,
		commands: map[string]scripts.ScriptFunc{
			"/start":         base.Start,
			"/create_task":   t.TaskCreation,
			"/tasks":         t.UserTasks,
			"/create_group":  g.GroupCreation,
			"/groups":        g.UserGroups,
			"/accept_invite": inv.InviteAcceptation,
			"/invites":       inv.UserInvites,
			"/invite":        inv.Invitation,
		},
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
	err := r.defineUser(update.Message.From)

	if err != nil {
		r.client.SendError(strconv.Itoa(update.Message.From.ID), "")
		return
	}

	s, sessionExists := r.sessions[update.Message.From.ID]

	if sessionExists {
		s.CancelCtx()
	}

	script, scriptExists := r.commands[update.Message.Text]

	if scriptExists {
		// создаем новую сессию
		session := NewSession(update.Message.From)

		// записываем для пользователя новую сессию
		r.sessions[update.Message.From.ID] = session

		handler := script(scripts.Script{
			Session: session,
			Storage: r.storage,
			Client:  r.client,
		})

		go func() {
			defer func() {
				session.CancelCtx()
				delete(r.sessions, update.Message.From.ID)

			}()
			err := handler.Run()

			if err != nil {
				log.Println(err)
				
				if errors.Is(err, e.ErrClientError) {
					r.client.SendError(strconv.Itoa(update.Message.From.ID), err.Error())
				}
				if errors.Is(err, e.ErrServerError) {
					r.client.SendError(strconv.Itoa(update.Message.From.ID), messages.InternalError)

				}
			}
		}()
	} else {
		r.client.SendError(strconv.Itoa(update.Message.From.ID), messages.CommandNotDefine)
	}
}
