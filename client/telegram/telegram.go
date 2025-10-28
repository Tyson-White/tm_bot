package telegram

type TelegramClient struct{}

func NewTelegramClient() TelegramClient {
	return TelegramClient{}
}

func (t TelegramClient) Register(apiBotToken string) error {

	return nil
}

func (t TelegramClient) Updates() {
	//TODO implement me
	panic("implement me")
}

func (t TelegramClient) SendMessage() {
	//TODO implement me
	panic("implement me")
}

func (t TelegramClient) SendPhoto() {
	//TODO implement me
	panic("implement me")
}
