package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

// Cleanup deletes the deployment using a function provided by the provisioner.
func Cleanup(configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var prv provisioner.Provisioner

	if configuration.Provisioner.Name == helm {
		cli = cmd.CLI{Binary: helm}
		prv = &provisioner.Helm{Logger: logger}
	}

	err = prv.Cleanup(configuration, &cli)

	return err
}
