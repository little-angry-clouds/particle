package driver

import (
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Driver interface {
	Create(config.ParticleConfiguration, cmd.Cmd) error
	Destroy(config.ParticleConfiguration, cmd.Cmd) error
}
