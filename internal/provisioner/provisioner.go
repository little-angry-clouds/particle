package provisioner

import (
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Provisioner interface {
	Converge(config.ParticleConfiguration, cmd.Cmd) error
	Cleanup(config.ParticleConfiguration, cmd.Cmd) error
	Dependency(config.ParticleConfiguration, cmd.Cmd) error
	Prepare(config.ParticleConfiguration, cmd.Cmd) error
}
