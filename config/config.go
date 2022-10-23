package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MAJOR_STRING_MATCHER                []string `envconfig:"MAJOR_STRING_MATCHER" default:"BREAKING CHANGE"`
	MINOR_STRING_MATCHER                []string `envconfig:"MINOR_STRING_MATCHER" default:"feat"`
	EXPORT_ENV_SEMVER_FULL_NAME         string   `envconfig:"EXPORT_ENV_SEMVER_FULL_NAME" default:"SEMVER_FULL"`
	EXPORT_ENV_SEMVER_MAJOR_NAME        string   `envconfig:"EXPORT_ENV_SEMVER_MAJOR_NAME" default:"SEMVER_MAJOR"`
	EXPORT_ENV_SEMVER_MINOR_NAME        string   `envconfig:"EXPORT_ENV_SEMVER_MINOR_NAME" default:"SEMVER_MINOR"`
	EXPORT_ENV_SEMVER_PATCH_NAME        string   `envconfig:"EXPORT_ENV_SEMVER_PATCH_NAME" default:"SEMVER_PATCH"`
	EXPORT_ENV_SEMVER_BUILD_NUMBER_NAME string   `envconfig:"EXPORT_ENV_SEMVER_BUILD_NUMBER_NAME" default:"SEMVER_BUILD_NUMBER"`
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
