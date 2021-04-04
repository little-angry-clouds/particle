package cli

import (
	"context"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

func Cleanup(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI
	var prv provisioner.Provisioner
	var ctx context.Context = context.Background()
	var prepare config.Key = "prepare"

	if configuration.Provisioner.Name == helm {
		cli = cmd.CLI{Binary: "helm"}
		prv = &provisioner.Helm{Logger: logger}
	}

	// Pass variables to context
	ctx = context.WithValue(ctx, prepare, configuration.Prepare)

	err = prv.Cleanup(ctx, &cli)

	return err
}
