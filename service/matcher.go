package service

import (
	"fmt"
	"regexp"

	"github.com/sirupsen/logrus"
)

type matcher struct {
	majorChangeMatcher []*regexp.Regexp
	minorChangeMatcher []*regexp.Regexp
}

type Matcher interface {
	IsMajorChange(value []byte) bool
	IsMinorChange(value []byte) bool
}

func NewMatcher(majorStringMatcher []string, minorStringMatcher []string) Matcher {
	majorChangeMatcher := make([]*regexp.Regexp, 0)
	minorChangeMatcher := make([]*regexp.Regexp, 0)

	majorStringMatcherSize := len(majorStringMatcher)
	minorStringMatcherSize := len(minorStringMatcher)
	if majorStringMatcherSize == 0 {
		logrus.Warn("The MajorChange will always return false.")
	}
	if minorStringMatcherSize == 0 {
		logrus.Warn("The MinorChange will always return false.")
	}

	// Generate MajorChange Matcher
	for i := 0; i < len(majorChangeMatcher); i++ {
		regMatcher, _ := newRegMatcher(majorStringMatcher[i])
		majorChangeMatcher = append(majorChangeMatcher, regMatcher)
	}

	// Generate MinorChange Matcher
	for i := 0; i < len(minorStringMatcher); i++ {
		regMatcher, _ := newRegMatcher(minorStringMatcher[i])
		minorChangeMatcher = append(minorChangeMatcher, regMatcher)
	}

	return &matcher{
		majorChangeMatcher: majorChangeMatcher,
		minorChangeMatcher: minorChangeMatcher,
	}
}

func newRegMatcher(value string) (*regexp.Regexp, error) {
	return regexp.Compile(fmt.Sprint(value, ".*:.*"))
}

func (m *matcher) IsMajorChange(value []byte) bool {
	for i := 0; i < len(m.majorChangeMatcher); i++ {
		if m.majorChangeMatcher[i].Match(value) {
			return true
		}
	}
	return false
}

func (m *matcher) IsMinorChange(value []byte) bool {
	for i := 0; i < len(m.minorChangeMatcher); i++ {
		if m.minorChangeMatcher[i].Match(value) {
			return true
		}
	}
	return false
}
