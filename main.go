package main

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func Exec(cmd string, shell bool) (string, error) {

	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		outStr := string(out)
		if err != nil {
			return outStr, err
		}
		return outStr, nil
	} else {
		out, err := exec.Command(cmd).Output()
		outStr := string(out)
		if err != nil {
			return outStr, err
		}
		return outStr, nil
	}
}

func main() {
	logrus.Info("Start determining the semver of this repo.")

	// Get the latest tag
	logrus.Info("Start getting the latest tag.")
	latestTag, err := Exec("git tag -l --sort=-creatordate | head -n 1", true)
	if err != nil {
		logrus.Error("Can't get latest tag")
		os.Exit(1)
		return
	}
	logrus.Debug(latestTag)
	if latestTag == "" {
		logrus.Warn("No latest tag was found.")
	}

	// Start determining sember.

}
