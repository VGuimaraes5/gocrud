package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type database struct {
	Host     string `envconfig:"DATABASE_HOST"`
	Port     string `envconfig:"DATABASE_PORT"`
	User     string `envconfig:"DATABASE_USER"`
	Password string `envconfig:"DATABASE_PASSWORD"`
	DBName   string `envconfig:"DATABASE_NAME"`
	SSLMode  string `envconfig:"DATABASE_SSL_MODE"`
	TimeZone string `envconfig:"DATABASE_TIMEZONE"`
}

func newDatabase() error {
	err := envconfig.Process("", &Database)
	if err != nil {
		return err
	}
	return nil
}
