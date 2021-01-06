package client

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Dial() error {
	return nil
}
