package app

type Config struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Addr: ":8080",
	}
}
