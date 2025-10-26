package models

import "fmt"

type TaskGroup struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}

func (t *TaskGroup) ToString() string {
	return fmt.Sprintf(`

👥 Группа: <b><u>%v</u></b>

ID <code>%v</code>
Создатель - @%v
`, t.Name, t.ID, t.Creator)
}
