package config

type Config struct {
	Server  Server  `env:"SERVER"`
	Logging Logging `env:"LOGGING"`
}

type Server struct {
	Host    string `env:"APP_HOST" required:"true" env-default:"localhost"`
	Port    string `env:"APP_PORT" required:"true" env-default:"8080"`
	GinMode string `env:"GIN_MODE" required:"true" env-default:"debug"`
}

type Logging struct {
	LogIndex  string `env:"LOG_INDEX" required:"true" env-default:"log"`
	IsDebug   bool   `env:"IS_DEBUG" required:"true" env-default:"true"`
	LogToFile bool   `env:"LOG_TO_FILE" required:"true" env-default:"false"`
}
