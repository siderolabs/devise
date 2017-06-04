package devise

import (
	"github.com/autonomy/devise/devise/backend"
	"github.com/autonomy/devise/devise/storage/datastore"
	"github.com/autonomy/devise/devise/ui"
)

// Server represents the Devise server.
type Server struct {
	Storage datastore.Datastore
}

// ServeOptions is used to configure the server.
type ServeOptions struct {
	Storage      string
	BackendPort  string
	UIPort       string
	VaultAddress string
}

// Start starts the server.
func Start(opts *ServeOptions) {
	go backend.Start(opts.BackendPort, opts.Storage, opts.VaultAddress)
	ui.Start(opts.UIPort)
}
