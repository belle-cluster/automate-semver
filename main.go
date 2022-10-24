package main

import (
	"fmt"
	"os"

	"github.com/flutter-semver/config"
	"github.com/flutter-semver/service"
	"github.com/sirupsen/logrus"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
	if appConfig.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	logrus.Info("Start determining the semver of this repo.")

	// Init Service
	executer := service.NewExecuter()
	git := service.NewGit(executer)
	matcher := service.NewMatcher(appConfig.MAJOR_STRING_MATCHER, appConfig.MINOR_STRING_MATCHER)

	// Get the latest tag
	latestTag, err := git.GetLatestTag()
	if err != nil {
		os.Exit(1)
		return
	}

	semver := service.NewSemverFromTag(string(latestTag))

	// Get Log
	logs, err := git.GetLog(string(latestTag), "HEAD")
	if err != nil {
		os.Exit(1)
		return
	}

	// Check if log fall into these condition
	if matcher.IsMajorChange(logs) {
		semver.BumpMajor()
	} else if matcher.IsMinorChange(logs) {
		semver.BumpMinor()
	} else {
		semver.BumpPatch()
	}
	semver.BumpBuildNumber()
	logrus.Info("Result: ", semver.Render())
	envManager := service.NewEnv()
	envManager.SetEnv(appConfig.EXPORT_ENV_SEMVER_FULL_NAME, semver.Render())
	envManager.SetEnv(appConfig.EXPORT_ENV_SEMVER_MAJOR_NAME, fmt.Sprint(semver.GetMajor()))
	envManager.SetEnv(appConfig.EXPORT_ENV_SEMVER_MINOR_NAME, fmt.Sprint(semver.GetMinor()))
	envManager.SetEnv(appConfig.EXPORT_ENV_SEMVER_PATCH_NAME, fmt.Sprint(semver.GetPatch()))
	envManager.SetEnv(appConfig.EXPORT_ENV_SEMVER_BUILD_NUMBER_NAME, fmt.Sprint(semver.GetBuildNumber()))
	envManager.WriteToFile()
}
