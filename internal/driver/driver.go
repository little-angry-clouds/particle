package driver

import (
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

// Driver is the interface that manages the kubernetes clusters.
// It always receives the particle configuration, a cmd.Cmd interface and
// returns error.
type Driver interface {
	Create(config.ParticleConfiguration, cmd.Cmd) error
	Destroy(config.ParticleConfiguration, cmd.Cmd) error
}
