package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

// Prepare installs the dependencies using a function provided from the provisioner.
func Prepare(configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var prv provisioner.Provisioner

	if configuration.Dependency.Name == helm {
		cli = cmd.CLI{Binary: "helm"}
		prv = &provisioner.Helm{Logger: logger}
	}

	err = prv.Prepare(configuration, &cli)

	return err
}
