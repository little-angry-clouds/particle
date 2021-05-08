package verifier

import (
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Verifier interface {
	Verify(config.ParticleConfiguration, cmd.Cmd) error
}
