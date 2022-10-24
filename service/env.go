package service

import (
	"fmt"
	"strings"
)

type env struct {
}

type Env interface {
	SetEnv(key string, value string)
}

func NewEnv() Env {
	return &env{}
}

func (e *env) SetEnv(key string, value string) {
	fmt.Printf("%s=%s\n", strings.ToUpper(key), value)
}
