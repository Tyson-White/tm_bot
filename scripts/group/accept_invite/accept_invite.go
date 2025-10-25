package accept_invite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"tg-bot/pkg/e"
	"tg-bot/pkg/models"
	"tg-bot/scripts"
	"tg-bot/types"
)

type Command struct{ types.ScriptInitParams }

func New(params types.ScriptInitParams) types.ScriptCommandHandler {
	return &Command{ScriptInitParams: params}
}

func (c *Command) Run() {

	invite, err := c.checkInvitation()

	if err != nil {
		log.Println(err)
		return
	}

	_, err = c.Storage.AddUserToGroup(c.Session.User.Username, invite.Group)

	if err != nil {
		c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), ErrGroupCreate)
		return
	}

	c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), fmt.Sprintf(`
	Теперь ты состоишь в группе %v
	`, invite.Group))

}

func (c *Command) checkInvitation() (models.Invite, error) {
	params := scripts.InputParams{
		Client:  c.Client,
		Session: c.Session,
		Msg:     GroupNameMSG,
	}

	upd, err := scripts.Input(params)
	if err != nil {
		return models.Invite{}, err
	}

	id, err := strconv.Atoi(upd.Message.Text)

	if err != nil {
		c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), ErrIdNotInt)
		return models.Invite{}, err
	}

	invite, err := c.Storage.InviteById(id, c.Session.User.Username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), ErrIdNotInt)
			return models.Invite{}, err
		}

		c.Client.SendMessage(strconv.Itoa(c.Session.User.ID), e.ErrServerMSG)
		return models.Invite{}, err
	}

	return invite, nil
}
