package config

type Config struct {
	Server   Server   `env:"SERVER"`
	Logging  Logging  `env:"LOGGING"`
	Postgres Postgres `env:"POSTGRES"`
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

type Postgres struct {
	Host            string `env:"POSTGRES_HOST" required:"true" env-default:"localhost"`
	Port            string `env:"POSTGRES_PORT" required:"true" env-default:"5432"`
	User            string `env:"POSTGRES_USER" required:"true" env-default:"postgres"`
	Password        string `env:"POSTGRES_PASSWORD" required:"true" env-default:"postgres"`
	Database        string `env:"POSTGRES_DB" required:"true" env-default:"postgres"`
	SslMode         string `env:"SSL_MODE" required:"true" env-default:"disable"`
	MaxConns        string `env:"MAX_CONNS" required:"true" env-default:"10"`
	ConnMaxLifetime string `env:"CONN_MAX_LIFETIME" required:"true" env-default:"10m"`
	ConnMaxIdleTime string `env:"CONN_MAX_IDLETIME" required:"true" env-default:"5m"`
}
