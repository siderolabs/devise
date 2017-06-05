package storage

import (
	"github.com/autonomy/devise/devise/storage/datastore"
	"github.com/autonomy/devise/devise/storage/datastore/memory"
)

var datastores = map[string]func() datastore.Datastore{
	"memory": func() datastore.Datastore { return memory.New() },
}

// New instantiates and returns a storage datastore
func New(b string) datastore.Datastore {
	return datastores[b]()
}
