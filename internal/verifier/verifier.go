package verifier

import (
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

// Driver is the interface that verifies that the state of what is
// deployed, is in the desired state.
type Verifier interface {
	Verify(config.ParticleConfiguration, cmd.Cmd) error
}
