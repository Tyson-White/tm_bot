package create_group

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tg-bot/pkg/models"
	"tg-bot/scripts"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (com *Command) Run() {
	params := scripts.InputParams{
		Client:    com.Client,
		Session:   com.Session,
		Msg:       GroupNameMSG,
		PhotoPath: "./assets/create_group.png",
	}

	nameUpd, err := scripts.Input(params)
	if err != nil {
		return
	}

	params.Msg = UsersMSG
	usersUpd, err := scripts.Input(params)
	if err != nil {
		return
	}

	group, err := com.Storage.CreateGroup(nameUpd.Message.Text, nameUpd.Message.From.Username)

	if err != nil {

		com.Client.SendFMessage(strconv.Itoa(com.Session.User.ID), GroupCreateErr)
		return
	}

	users := strings.Split(usersUpd.Message.Text, " ")

	failedInvites := []string{}
	invited := []string{}

	for _, us := range users {
		if strings.HasPrefix(us, "@") {
			// delete this prefix
			err := com.Storage.CreateInvite(group.Name, com.Session.User.Username, us[1:])

			if err != nil {
				failedInvites = append(failedInvites, us)
			} else {
				invited = append(invited, us)
				com.sendInviteMessage(us[1:], group)
			}
		}
	}

	com.Client.SendFMessage(strconv.Itoa(com.Session.User.ID), fmt.Sprintf(`
Создана группа
%v Приглашены: %v
`, group.ToString(), invited, failedInvites))

	// TODO: Отправлять приглашения пользователям

}

func (com *Command) sendInviteMessage(username string, group models.TaskGroup) {
	user, err := com.Storage.UserByName(username)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(user)

	com.Client.SendFMessage(strconv.Itoa(user.Telegram), fmt.Sprintf(`
		Вас пригласили в группу

		%v
	`, group.ToString()))
}
