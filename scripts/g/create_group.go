package g

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tg-bot/pkg/messages"
	"tg-bot/pkg/models"
	"tg-bot/scripts"
)

type GroupCreationCMD struct{ scripts.Script }

func GroupCreation(params scripts.Script) scripts.ScriptMethods {
	return &GroupCreationCMD{Script: params}
}

func (cmd *GroupCreationCMD) Run() error {

	photo := "./assets/create_group.png"

	nameUpd, err := cmd.RequestInput(messages.ReqGroupName, photo)
	if err != nil {
		return err
	}

	usersUpd, err := cmd.RequestInput(messages.UsersMSG, photo)
	if err != nil {
		return err
	}

	group, err := cmd.Storage.CreateGroup(nameUpd.Message.Text, nameUpd.Message.From.Username)

	if err != nil {

		cmd.Err(messages.GroupCreateErr)
		return err
	}

	if usersUpd.Message.Text != "skip" {
		users := strings.Split(usersUpd.Message.Text, " ")

		cmd.sendInvites(group, users)
	}

	cmd.Success(fmt.Sprintf(`Создана группа
%v
`, group.ToString()))

	return nil
}

func (cmd *GroupCreationCMD) sendInvites(group models.TaskGroup, userNames []string) {
	for _, us := range userNames {
		if strings.HasPrefix(us, "@") {
			_, err := cmd.Storage.CreateInvite(group.Name, cmd.Session.User.Username, us[1:])

			if err == nil {
				cmd.sendInviteMessage(us[1:], group)
			}
		}
	}
}

func (cmd *GroupCreationCMD) sendInviteMessage(username string, group models.TaskGroup) {
	user, err := cmd.Storage.UserByName(username)

	if err != nil {
		log.Println(err)
		return
	}

	cmd.Client.SendFMessage(strconv.Itoa(user.Telegram), fmt.Sprintf(`
		Вас пригласили в группу

		%v
	`, group.ToString()))
}
