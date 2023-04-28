package config

var Default = &Config{
	Addr:  "0.0.0.0:8080",
	Route: map[string]string{},
}

type Config struct {
	Route    map[string]string `json:"route"`
	Addr     string            `json:"addr"`
	NotFound string            `json:"not_found"`
}

func (c *Config) Merge(c2 *Config) {
	for k, v := range c2.Route {
		c.Route[k] = v
	}

	if c2.Addr != "" {
		c.Addr = c2.Addr
	}

	c.NotFound = c2.NotFound
}

func (c *Config) GetRoute() map[string]string {
	return c.Route
}

func (c *Config) GetAddr() string {
	return c.Addr
}
