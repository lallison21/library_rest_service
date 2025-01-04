package config

type Config struct {
	Server Server `env:"SERVER"`
}

type Server struct {
	Host    string `env:"APP_HOST" env-default:"localhost"`
	Port    string `env:"APP_PORT" env-default:"8080"`
	GinMode string `env:"GIN_MODE" env-default:"debug"`
}
