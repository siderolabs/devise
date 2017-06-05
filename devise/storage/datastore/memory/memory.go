package memory

import (
	"github.com/autonomy/devise/devise/storage/datastore"
)

// Memory represents a datastore
type Memory struct {
	store map[string][]byte
}

// New instantiates and returns a datastore.Datastore.
func New() datastore.Datastore {
	return Memory{store: make(map[string][]byte)}
}

// Get implements datastore.Datastore.
func (m Memory) Get(k string) (*datastore.Entry, error) {
	if value, ok := m.store[k]; ok {
		return &datastore.Entry{Key: k, Value: value}, nil
	}

	return &datastore.Entry{Key: k, Value: nil}, nil
}

// Put implements datastore.Datastore.
func (m Memory) Put(e *datastore.Entry) error {
	m.store[e.Key] = e.Value

	return nil
}

// Delete implements datastore.Datastore.
func (m Memory) Delete(k string) error {
	delete(m.store, k)

	return nil
}
