package models

import "fmt"

type Invite struct {
	ID      int    `json:"id" db:"id"`
	Group   string `json:"groupname" db:"groupname"`
	Creator string `json:"creator" db:"creator"`
	Invited string `json:"invited" db:"invited"`
}

func (inv *Invite) ToString() string {

	return fmt.Sprintf(`
		Приглашение %d
		Группа: %d
		Пригласивший: @%v
	`, inv.ID, inv.Group, inv.Invited)
}
