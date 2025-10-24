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
	return fmt.Sprintf(`
	Задача: %v
	
	Описание: 
	%v

	🔗 Создал: @%v

	🕝 Создана %v
	
	`,
		t.Title,
		t.Description,
		t.Owner,
		t.CreatedAt)
}
