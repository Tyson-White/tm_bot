package client

type Client struct{}

func New() Client {
	return Client{}
}

func (c Client) Register(apiBotToken string) error {
	return nil
}

func (c Client) Updates() {
	//TODO implement me
	panic("implement me")
}
