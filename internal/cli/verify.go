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

	if configuration.Verifier.Name == "helm" {
		cli = cmd.CLI{Binary: "helm"}
		vrf = &verifier.Helm{Logger: logger}
	}

	err = vrf.Verify(configuration, &cli)

	return err
}
