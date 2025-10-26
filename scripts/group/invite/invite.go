package invite

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tg-bot/scripts"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (cmd *Command) Run() {

	p := scripts.InputParams{
		Client:  cmd.Client,
		Session: cmd.Session,
		Msg:     inviteGroupMSG,
	}

	upd, err := scripts.Input(p)

	if err != nil {
		return
	}

	params := strings.Split(upd.Message.Text, " ")

	if len(params) < 2 {
		cmd.Client.SendError(strconv.Itoa(cmd.Session.User.ID), "🚫 Неверный формат")
		return
	}

	success := []string{}

	for _, us := range params[1:] {
		if !strings.HasPrefix(params[1], "@") {
			cmd.Client.SendError(strconv.Itoa(cmd.Session.User.ID), "🚫 Неверный формат")
			continue
		}

		err := cmd.Storage.CreateInvite(params[0], cmd.Session.User.Username, us[1:])

		if err != nil {
			log.Println(err)
			continue
		}

		success = append(success, us[1:])
	}

	cmd.Client.SendSuccess(strconv.Itoa(cmd.Session.User.ID), "✅ Приглашения отправлены")

	for _, us := range success {
		user, err := cmd.Storage.UserByName(us)
		log.Println(us)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(user.ID)

		msg := fmt.Sprintf(`
Тебя пригласили в группу %v
Принять приглашение /accept_invite
		`, params[0])

		cmd.Client.SendFMessage(strconv.Itoa(user.Telegram), msg)
	}
}
