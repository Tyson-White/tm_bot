package models

import "fmt"

type TaskGroup struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

func (t *TaskGroup) ToString() string {
	return fmt.Sprintf(`
	Группа: %d
	Название: %v

	Создатель - %v
	`, t.ID, t.Name, t.Creator)
}
