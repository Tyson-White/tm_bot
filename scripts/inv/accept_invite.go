package inv

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"tg-bot/pkg/e"
	"tg-bot/pkg/messages"
	"tg-bot/pkg/models"
	"tg-bot/scripts"
)

type InviteAcceptationCMD struct{ scripts.Script }

func InviteAcceptation(params scripts.Script) scripts.ScriptMethods {
	return &InviteAcceptationCMD{Script: params}
}

func (cmd *InviteAcceptationCMD) Run() error {

	invite, err := cmd.checkInvitation()

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = cmd.Storage.AddUserToGroup(cmd.Session.User.Username, invite.Group)

	if err != nil {
		cmd.Err(messages.ErrGroupCreate)
		return err
	}

	cmd.Msg(fmt.Sprintf(`
	Теперь ты состоишь в группе %v
	`, invite.Group), "")

	return nil
}

func (cmd *InviteAcceptationCMD) checkInvitation() (models.Invite, error) {

	upd, err := cmd.RequestInput(messages.GroupNameMSG, "")
	if err != nil {
		return models.Invite{}, err
	}

	id, err := strconv.Atoi(upd.Message.Text)

	if err != nil {
		cmd.Err(messages.ErrIdNotInt)
		return models.Invite{}, err
	}

	invite, err := cmd.Storage.InviteById(id, cmd.Session.User.Username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cmd.Err(messages.ErrIdNotInt)
			return models.Invite{}, err
		}

		cmd.Err(e.ErrServerMSG)
		return models.Invite{}, err
	}

	return invite, nil
}
