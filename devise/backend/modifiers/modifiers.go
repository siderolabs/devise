package modifiers

// Modifiers is an object used to hold all modifiers for use within templates.
type Modifiers struct {
	Env   *Env
	Vault *Vault
}

// Options holds all the modififer specific settings.
type Options struct {
	VaultAddress string
}

// NewModifiers instantiates and returns a Modifiers.
func NewModifiers(opts *Options) *Modifiers {
	return &Modifiers{
		Env:   NewEnv(),
		Vault: NewVault(opts.VaultAddress),
	}
}
