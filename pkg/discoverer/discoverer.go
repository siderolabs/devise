package discoverer

import (
	"github.com/autonomy/devise/pkg/discoverer/kubernetes"
	"github.com/autonomy/devise/pkg/renderer"
	"github.com/autonomy/devise/pkg/storage/datastore"
)

var discoverersMap = map[string]func() Discoverer{
	"kubernetes": func() Discoverer { return kubernetes.New() },
}

// Discoverer is an interface for finding Devise clients.
type Discoverer interface {
	Discover(datastore.Datastore, *renderer.Renderer)
}

// NewDiscoverers instantiates and returns a slice of discoverers.
func NewDiscoverers(d []string) (discoverers []Discoverer) {
	for _, e := range d {
		discoverers = append(discoverers, discoverersMap[e]())
	}

	return
}
