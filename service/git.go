package service

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type git struct {
	executer Executer
}

type Git interface {
	GetLatestTag() ([]byte, error)
	GetLog(from string, to string) ([]byte, error)
}

func NewGit(executer Executer) Git {
	return &git{
		executer: executer,
	}
}

func (g *git) GetLatestTag() ([]byte, error) {
	logrus.Info("Start getting the latest tag.")
	latestTag, err := g.executer.Exec("git tag -l --sort=-creatordate | head -n 1", true)
	logrus.Info("Current Tag: ", string(latestTag))
	if err != nil {
		logrus.Error(err)
	}

	if len(latestTag) == 0 {
		logrus.Warn("No latest tag was found.")
	}

	return latestTag, err
}

func (g *git) GetLog(from string, to string) ([]byte, error) {
	command := "git log --pretty=oneline"
	if from != "" {
		command = fmt.Sprintf("%s %s..%s", command, from, to)
	}
	logrus.Debug("Execution command: ", command)
	log, err := g.executer.Exec(command, true)
	logrus.Debug(string(log))
	if err != nil {
		logrus.Error(err)
	}

	if len(log) == 0 {
		logrus.Warn("No history found.")
	}
	return log, err
}
