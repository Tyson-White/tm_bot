package client

type Sender struct {
	chatId int64
	client Client
}

func NewSender(chatId int64, client Client) Sender {
	return Sender{
		chatId: chatId,
		client: client,
	}
}

func (s *Sender) SendMsg(msg string) {
	//s.client.SendMessage()
}
