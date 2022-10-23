package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/sirupsen/logrus"
)

func Exec(cmd string, shell bool) ([]byte, error) {

	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return out, err
		}
		return out, nil
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			return out, err
		}
		return out, nil
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
	if len(latestTag) == 0 {
		logrus.Warn("No latest tag was found.")
	}

	// Start determining semver.

	var diffResult []byte
	// Have latestTag.
	if len(latestTag) != 0 {
		diffResult, err = Exec(fmt.Sprintf("git log --pretty=oneline %s..HEAD", latestTag), true)
	} else {
		diffResult, err = Exec("git log --pretty=oneline", true)
	}

	if err != nil {
		logrus.Error("Get diff failed.")
		os.Exit(1)
		return
	}

	if len(diffResult) == 0 {
		logrus.Warn("Diff Result is empty")
	}

	// Match semver
	if matched, _ := regexp.Match("BREAKING CHANGE:.*", diffResult); matched {
		// TODO: Bump major verison up.
	} else if matched, _ := regexp.Match("feat:.*", diffResult); matched {
		// TODO: Bump minor version up.
	} else {
		// TODO: Bump patch version up.
	}

	// TODO: Bump build number up.

}
