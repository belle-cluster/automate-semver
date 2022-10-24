package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type env struct {
	data map[string]string
}

type Env interface {
	SetEnv(key string, value string)
	WriteToFile()
}

func NewEnv() Env {
	return &env{
		data: make(map[string]string),
	}
}

func (e *env) SetEnv(key string, value string) {
	fmt.Printf("%s=%s\n", strings.ToUpper(key), value)
	e.data[key] = value
}

func (e *env) WriteToFile() {
	f, err := os.Create("semver-result.txt")
	if err != nil {
		logrus.Error(err)
		return
	}
	for key, value := range e.data {
		f.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	}
}
