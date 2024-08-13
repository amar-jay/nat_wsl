package portmap

import (
	"fmt"
	"sync"

	"github.com/amar-jay/nat_wsl/pkg/config"
)

type PortMaps struct {
	*IPTable
}

/**
* NewPortMaps is a service that maps the configuration file to a map of port mappings.
 */
func NewPortMaps(c *config.Config) (*PortMaps, error) {
	if c == nil {
		return nil, fmt.Errorf("config is empty")
	}

	cache := NewIPTable()
	for k, v := range *c {
		if v.Type != "v4tov4" {
			return nil, fmt.Errorf("Unsupported type: %s", v.Type)
		}

		if v.Wsl.Listenip == "" {
			v.Wsl.Listenip = "localhost"
		}

		_pm := PortMapping{
			Protocol:     v.Protocol,
			Externalip:   v.Wsl.Listenip,
			Externalport: v.Wsl.Listenport,
			Internalip:   v.Remote.Connectip,
			Internalport: v.Remote.Connectport,
		}

		cache.AddMapping(k, _pm)
	}
	return &PortMaps{cache}, nil
}

// start starts all port mapping service.
func (pm PortMaps) Start() {
	if pm.IPTable == nil {
		panic(fmt.Errorf("Ip table not initialized"))
	}

	var wg sync.WaitGroup
	for _, v := range pm.ListMappings() {
		wg.Add(1)
		go v.Start() // TODO: make this more optimized, too many threads running...
	}
	wg.Wait()
}
