package telegram

import "net/http"

type Client struct {
	protocol string
	host     string
	baseUrl  string
	client   http.Client
}

type SendMessageBody struct {
	Chat string `json:"chat_id"`
	Text string `json:"text"`
	Mode string `json:"parse_mode"`
}

type UpdateEntity struct {
	ID      int           `json:"update_id"`
	Message MessageEntity `json:"message"`
}

type MessageEntity struct {
	ID   int        `json:"message_id"`
	From UserEntity `json:"from"`
	Text string     `json:"text"`
	Date int        `json:"date"`
}

type UserEntity struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UpdatesResponse struct {
	Ok     bool           `json:"ok"`
	Result []UpdateEntity `json:"result"`
}
