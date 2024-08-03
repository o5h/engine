package app

type Config struct {
	Width, Height int32
	Title         string
}

var DefaultConfig = Config{
	Width:  640,
	Height: 480,
	Title:  "Example",
}

func WithTitle(name string) func(*Config) { return func(c *Config) { c.Title = name } }
