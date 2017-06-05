package modifiers

import "github.com/hashicorp/vault/api"

// Vault represents the vault modifier.
type Vault struct {
	Address string
	Client  *api.Client
}

// NewVault instantiates and returns a reference to a Vault
func NewVault(a string) *Vault {
	client, _ := api.NewClient(&api.Config{Address: a})
	return &Vault{
		Address: a,
		Client:  client,
	}
}

// Secret returns the secret.
func (v *Vault) Secret(p, k string) (string, error) {
	secret, err := v.Client.Logical().Read(p)
	if err != nil {
		return "nil", err
	}

	return secret.Data[k].(string), nil
}
