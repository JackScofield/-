package config

type config struct {
	natsURL string
	port    string
}

func (c *config) Port() string {
	return c.port
}

func (c *config) SetPort(port string) {
	c.port = port
}

func (c *config) NatsURL() string {
	return c.natsURL
}

func (c *config) SetNatsURL(natsURL string) {
	c.natsURL = natsURL
}
