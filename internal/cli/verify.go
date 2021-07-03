package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/verifier"
)

func Verify(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var vrf verifier.Verifier

	cli = cmd.CLI{Binary: configuration.Verifier.Command[0]}
	vrf = &verifier.Command{Logger: logger}

	err = vrf.Verify(configuration, &cli)

	return err
}
