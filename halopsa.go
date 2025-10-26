package halopsa

import (
	"github.com/HarpreetVaid/GO-HaloPsa/client"
	"github.com/HarpreetVaid/GO-HaloPsa/tickets"
)

type Halo struct {
	Client  *client.Client
	Tickets *tickets.Service
}

func New(baseURL, token string) *Halo {
	c := client.NewClient(baseURL, token)
	return &Halo{
		Client:  c,
		Tickets: tickets.NewService(c),
	}
}
