package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/verifier"
)

// Verify checks that what's deployed on kubernetes has the desired state.
func Verify(configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var vrf verifier.Verifier
	var cli cmd.CLI = cmd.CLI{Binary: "bash"}

	vrf = &verifier.Command{Logger: logger}

	err = vrf.Verify(configuration, &cli)

	return err
}
