package env

import (
	"fmt"
	"os"
)

// Env represents the environment modifier.
type Env struct{}

// New instantiates and returns a reference to an Env
func New() *Env {
	return &Env{}
}

// Get returns the environment variable.
func (e *Env) Get(s string) (*string, error) {
	envvar := os.Getenv(s)
	if envvar == "" {
		return nil, fmt.Errorf("`%s` is missing", s)
	}

	return &envvar, nil
}
