package config

import "time"

type Config struct {
	Server   Server   `env:"SERVER"`
	Logging  Logging  `env:"LOGGING"`
	Postgres Postgres `env:"POSTGRES"`
	Password Password `env:"PASSWORD"`
}

type Server struct {
	Host              string        `env:"APP_HOST"            env-default:"localhost" required:"true"`
	Port              string        `env:"APP_PORT"            env-default:"8080"      required:"true"`
	GinMode           string        `env:"GIN_MODE"            env-default:"debug"     required:"true"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" env-default:"10s"       required:"true"`
}

type Logging struct {
	LogIndex  string `env:"LOG_INDEX"   env-default:"log"   required:"true"`
	IsDebug   bool   `env:"IS_DEBUG"    env-default:"true"  required:"true"`
	LogToFile bool   `env:"LOG_TO_FILE" env-default:"false" required:"true"`
}

type Postgres struct {
	Host            string        `env:"POSTGRES_HOST"     env-default:"localhost" required:"true"`
	Port            string        `env:"POSTGRES_PORT"     env-default:"5432"      required:"true"`
	User            string        `env:"POSTGRES_USER"     env-default:"postgres"  required:"true"`
	Password        string        `env:"POSTGRES_PASSWORD" env-default:"postgres"  required:"true"`
	Database        string        `env:"POSTGRES_DB"       env-default:"postgres"  required:"true"`
	SslMode         string        `env:"SSL_MODE"          env-default:"disable"   required:"true"`
	MaxConns        int32         `env:"MAX_CONNS"         env-default:"10"        required:"true"`
	ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME" env-default:"10m"       required:"true"`
	ConnMaxIdleTime time.Duration `env:"CONN_MAX_IDLETIME" env-default:"5m"        required:"true"`
}

type Password struct {
	Memory      uint32 `env:"PASSWORD_MEMORY"      env-default:"65536" required:"true"`
	Iterations  uint32 `env:"PASSWORD_ITERATIONS"  env-default:"1"     required:"true"`
	Parallelism uint8  `env:"PASSWORD_PARALLELISM" env-default:"4"     required:"true"`
	SaltLength  uint32 `env:"PASSWORD_SALT_LENGTH" env-default:"16"    required:"true"`
	KeyLength   uint32 `env:"PASSWORD_KEY_LENGTH"  env-default:"32"    required:"true"`
}
