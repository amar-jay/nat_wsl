package portmap

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

// IPTable represents a table of port mappings.
type IPTable struct {
	mappings *cache.Cache
}

// NewIPTable creates a new IP table.
// default expiration time is 600 hours and cleanup interval is 1000 hours.
func NewIPTable() *IPTable {
	return &IPTable{
		mappings: cache.New(600*time.Hour, 1000*time.Hour),
	}
}

// AddMapping adds a new port mapping to the table.
func (nt *IPTable) AddMapping(key string, mapping PortMapping) {
	//key := fmt.Sprintf("%s:%d", mapping.Internalip, mapping.Internalport)
	nt.mappings.Set(key, mapping, cache.DefaultExpiration)
}

// RemoveMapping removes a port mapping from the table.
func (nt *IPTable) RemoveMapping(internalIP string, internalPort int) {
	key := fmt.Sprintf("%s:%d", internalIP, internalPort)
	nt.mappings.Delete(key)
}

// GetMapping retrieves a port mapping from the table.
func (nt *IPTable) GetMapping(internalIP string, internalPort int) (PortMapping, bool) {
	key := fmt.Sprintf("%s:%d", internalIP, internalPort)
	value, found := nt.mappings.Get(key)
	if !found {
		return PortMapping{}, false
	}
	return value.(PortMapping), true
}

// ListMappings returns all port mappings in the table.
func (nt *IPTable) ListMappings() []PortMapping {
	items := nt.mappings.Items()
	mappings := make([]PortMapping, 0, len(items))
	for _, item := range items {
		mappings = append(mappings, item.Object.(PortMapping))
	}
	return mappings
}

// refreshMapping refreshes the expiration time of a port mapping.
func (nt *IPTable) refreshMapping() {
	nt.mappings.DeleteExpired()
}
