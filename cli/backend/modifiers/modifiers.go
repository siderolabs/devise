package modifiers

// Modifiers is an object used to hold all modifiers for use within templates.
type Modifiers struct {
	Env *Env
}

// NewModifiers instantiates and returns a Modifiers.
func NewModifiers() *Modifiers {
	return &Modifiers{
		Env: NewEnv(),
	}
}
