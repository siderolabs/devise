package modifiers

import (
	"fmt"
	"os"
)

type Env struct{}

func NewEnv() *Env {
	e := Env{}

	return &e
}

func (e *Env) Get(s string) (*string, error) {
	envvar := os.Getenv(s)
	if envvar == "" {
		return nil, fmt.Errorf("`%s` is missing", s)
	}

	return &envvar, nil
}
