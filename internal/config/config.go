package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP struct {
		IP                 string        `env:"HTTP_IP"`
		Port               uint16        `env:"HTTP_PORT"`
		GinMode            string        `env:"GIN_MODE"`
		ReadTimeout        time.Duration `env:"HTTP_READ_TIMEOUT"`
		WriteTimeout       time.Duration `env:"HTTP_WRITE_TIMEOUT"`
		IdleTimeout        time.Duration `env:"HTTP_IDLE_TIMEOUT"`
		MaxHeaderMebibytes int           `env:"HTTP_MAX_HEADER_MEBIBYTES"`
		ShutdownTimeout    time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT"`
	}

	PostgreSQL struct {
		Username string `env:"PSQL_USERNAME"`
		Password string `env:"PSQL_PASSWORD"`
		Host     string `env:"PSQL_HOST"`
		Port     uint16 `env:"PSQL_PORT"`
		DB       string `env:"PSQL_DB"`
	}

	PasswordSalt string `env:"PASSWORD_SALT"`

	Auth struct {
		AccessTokenTTL       time.Duration `env:"ACCESS_TOKEN_TTL"`
		RefreshTokenTTL      time.Duration `env:"REFRESH_TOKEN_TTL"`
		JwtSecretKey         string        `env:"JWT_SECRET_KEY"`
		MaxUserSessionsCount int64         `env:"MAX_USER_SESSIONS_COUNT"`
	}
}

func ReadEnvFile() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)

	return &cfg, err
}
