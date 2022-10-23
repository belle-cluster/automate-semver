package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MAJOR_STRING_MATCHER []string `envconfig:"MAJOR_STRING_MATCHER" default:"BREAKING CHANGE"`
	MINOR_STRING_MATCHER []string `envconfig:"MINOR_STRING_MATCHER" default:"feat"`
}

func Load() Config {
	var config Config

	err := godotenv.Load("./.env")
	if err != nil {
		logrus.Debug("No env file.")
	}

	envconfig.MustProcess("", &config)
	return config
}
