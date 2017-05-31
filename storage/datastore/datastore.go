package datastore

// Datastore represents a storage engine.
type Datastore interface {
	Get(string) (*Entry, error)
	Put(*Entry) error
	Delete(string) error
}

// Entry is the entry for an item in a Store implementation.
type Entry struct {
	Key   string
	Value []byte
}
