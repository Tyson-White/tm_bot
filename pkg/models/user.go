package models

type TUser struct {
	ID       int    `json:"id" bd:"id"`
	Telegram int    `json:"telegram_id" db:"telegram_id"`
	Username string `json:"username" db:"username"`
}
