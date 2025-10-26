package models

import (
	"fmt"
)

type Task struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	Group       *string `json:"groupname" db:"groupname"`
	Owner       string  `json:"owner" db:"owner"`
}

func (t *Task) ToString() string {

	group := "Нет"

	if t.Group != nil {
		group = *t.Group
	}
	return fmt.Sprintf(`
🗃️ Задача: <b><u>%v</u></b>
<blockquote expandable><b>Описание</b>
%v</blockquote>
ID <code>%v</code>
Группа: %v
Создатель: @%v`,
		t.Title,
		t.Description,
		t.ID,
		group,
		t.Owner)
}
