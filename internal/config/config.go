package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		IP      string `env:"HTTP_IP"`
		Port    uint16 `env:"HTTP_PORT"`
		GinMode string `env:"GIN_MODE"`
	}

	PostgreSQL struct {
		Username string `env:"PSQL_USERNAME"`
		Password string `env:"PSQL_PASSWORD"`
		Host     string `env:"PSQL_HOST"`
		Port     uint16 `env:"PSQL_PORT"`
		DB       string `env:"PSQL_DB"`
	}
}

func ReadEnvFile() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)

	return &cfg, err
}
