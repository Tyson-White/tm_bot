package inv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tg-bot/pkg/messages"
	"tg-bot/scripts"
)

type InvitationCMD struct{ scripts.Script }

func Invitation(params scripts.Script) scripts.ScriptMethods {
	return &InvitationCMD{Script: params}
}

type Invite struct {
	username string
	inviteId int
}

func (cmd *InvitationCMD) Run() error {

	upd, err := cmd.RequestInput(messages.InviteGroupMSG, "")

	if err != nil {
		return err
	}

	params := strings.Split(upd.Message.Text, " ")

	if len(params) < 2 {
		cmd.Msg("🚫 Неверный формат", "")
		return nil
	}

	success := make([]Invite, 0)

	for _, us := range params[1:] {
		if !strings.HasPrefix(us, "@") {
			cmd.Err("🚫 Неверный формат")
			continue
		}

		if us[1:] == cmd.Session.User.Username {
			cmd.Err("<b><u>Нельзя отправлять приглашения самому себе</u></b>")
			continue
		}

		id, err := cmd.Storage.CreateInvite(params[0], cmd.Session.User.Username, us[1:])

		if err != nil {
			log.Println(err)
			continue
		}

		success = append(success, Invite{
			username: us[1:],
			inviteId: id,
		})
	}

	if len(success) > 0 {
		cmd.Success("<b><u>Приглашения отправлены</u></b>")
	}

	for _, us := range success {
		invitedUs, err := cmd.Storage.UserByName(us.username)

		if err != nil {
			log.Println(err)
			continue
		}

		msg := fmt.Sprintf(`
Тебя пригласили в группу %v (ID приглашения: <code>%v</code>)
Принять приглашение /accept_invite
		`, params[0], us.inviteId)

		cmd.Client.SendFMessage(strconv.Itoa(invitedUs.Telegram), msg)
	}

	return nil
}
