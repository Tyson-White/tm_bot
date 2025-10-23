package models

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	GroupId     *int   `json:"group_id" db:"group_id"`
	Owner       string `json:"owner"`
}

func (t *Task) ToString() string {
	return fmt.Sprintf(`
	Задача: %v
	
	Описание: 
	%v

	🔗 Создал: %v

	🕝 Создана %v
	
	`,
		t.Title,
		t.Description,
		t.Owner,
		t.CreatedAt)
}
