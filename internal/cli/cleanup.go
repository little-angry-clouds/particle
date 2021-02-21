package cli

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

func Cleanup(scenario string, configuration config.ParticleConfiguration) error {
	var err error
	var cli cmd.CLI
	var prv provisioner.Provisioner
	var ctx context.Context = context.Background()

	configuration, err = config.ReadConfiguration(scenario)
	if err != nil {
		return err
	}

	if configuration.Provisioner.Name == helm {
		cli = cmd.CLI{Binary: "helm"}
		prv = &provisioner.Helm{}
	}

	err = prv.Cleanup(ctx, &cli)
	if err != nil {
		return err
	}

	return err
}
