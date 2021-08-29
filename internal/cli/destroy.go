package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/driver"
)

// Destroy creates the kubernetes cluster using a function provided by the driver.
func Destroy(configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var drv driver.Driver

	if configuration.Driver.Name == kind {
		cli = cmd.CLI{Binary: kind}
		drv = &driver.Kind{Logger: logger}
	}

	err = drv.Destroy(configuration, &cli)

	return err
}
