package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/verifier"
)

func Verify(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var vrf verifier.Verifier
	var cli cmd.CLI = cmd.CLI{Binary: "bash"}

	vrf = &verifier.Command{Logger: logger}

	err = vrf.Verify(configuration, &cli)

	return err
}
