package main

import (
	"os"

	"github.com/flutter-semver/config"
	"github.com/flutter-semver/service"
	"github.com/sirupsen/logrus"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
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

	// Get Log
	logs, err := git.GetLog(string(latestTag), "HEAD")
	if err != nil {
		os.Exit(1)
		return
	}

	// Check if log fall into these condition
	if matcher.IsMajorChange(logs) {
		// TODO: Bump Major Version Up
	} else if matcher.IsMinorChange(logs) {
		// TODO: Bump Minor Version Up
	} else {
		// TODO: Bump Patch Version Up
	}

	// TODO: Bump build number up

}
