package service

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

type semver struct {
	Major       int
	Minor       int
	Patch       int
	BuildNumber int
}

type Semver interface {
	Render() string
	BumpMajor()
	BumpMinor()
	BumpPatch()
	BumpBuildNumber()
	GetMajor() int
	GetMinor() int
	GetPatch() int
	GetBuildNumber() int
}

func NewSemver(major, minor, patch, buildNumber int) Semver {
	return &semver{
		Major:       major,
		Minor:       minor,
		Patch:       patch,
		BuildNumber: buildNumber,
	}
}

func NewSemverFromTag(tag string) Semver {
	tagRegexp, _ := regexp.Compile(`(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)
	result := tagRegexp.FindAllStringSubmatch(tag, -1)

	major := 0
	minor := 0
	patch := 0
	buildNumber := 0

	if len(result) == 0 {
		logrus.Warn("The inputted tag can't be matched by tagRegexp")
	} else {
		major, _ = strconv.Atoi(result[0][1])
		minor, _ = strconv.Atoi(result[0][2])
		patch, _ = strconv.Atoi(result[0][3])
		buildNumber, _ = strconv.Atoi(result[0][5])
	}

	return NewSemver(major, minor, patch, buildNumber)
}

func (s *semver) BumpMajor() {
	s.Major++
	s.Minor = 0
	s.Patch = 0
}

func (s *semver) BumpMinor() {
	s.Minor++
	s.Patch = 0
}

func (s *semver) BumpPatch() {
	s.Patch++
}

func (s *semver) BumpBuildNumber() {
	s.BuildNumber++
}

func (s *semver) GetMajor() int {
	return s.Major
}

func (s *semver) GetMinor() int {
	return s.Minor
}

func (s *semver) GetPatch() int {
	return s.Patch
}

func (s *semver) GetBuildNumber() int {
	return s.BuildNumber
}

func (s *semver) Render() string {
	return fmt.Sprintf("%d.%d.%d+%d", s.Major, s.Minor, s.Patch, s.BuildNumber)
}
