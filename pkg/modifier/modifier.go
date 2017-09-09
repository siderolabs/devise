package modifier

import (
	"github.com/autonomy/devise/pkg/modifier/env"
	"github.com/autonomy/devise/pkg/modifier/vault"
)

// Modifiers is an object used to hold all modifiers for use within templates.
type Modifiers struct {
	Env   *env.Env
	Vault *vault.Vault
}

// Options holds all the modififer specific settings.
type Options struct {
	VaultAddress string
}

// NewModifiers instantiates and returns a Modifiers.
func NewModifiers(opts *Options) *Modifiers {
	return &Modifiers{
		Env:   env.New(),
		Vault: vault.New(opts.VaultAddress),
	}
}
