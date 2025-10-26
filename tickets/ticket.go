package tickets

import "github.com/HarpreetVaid/GO-HaloPsa/client"

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{client: c}
}
