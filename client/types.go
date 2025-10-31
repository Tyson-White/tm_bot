package client

type Connection struct {
}

type Update struct {
	UpdateID int64 `json:"update_id"`

	Message *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int64 `json:"message_id"`
	From      User  `json:"from,omitempty"`
	Date      int64 `json:"date"`
	Chat      Chat  `json:"chat"`

	Text string `json:"text,omitempty"`
}

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	IsPremium bool   `json:"is_premium,omitempty"`
}

type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
