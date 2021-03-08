package cli

import (
	"context"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/driver"
)

func Create(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var drv driver.Driver
	var ctx context.Context = context.Background()
	var kubernetesVersion config.Key = "kubernetesVersion"

	if configuration.Driver.Name == "kind" {
		cli = cmd.CLI{Binary: "kind"}
		drv = &driver.Kind{Logger: logger}
	}

	// Pass variables to context
	if configuration.Driver.KubernetesVersion != "" {
		ctx = context.WithValue(ctx, kubernetesVersion, configuration.Driver.KubernetesVersion)
	}

	err = drv.Create(ctx, &cli)

	return err
}
