package cli

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

func Converge(scenario string, configuration config.ParticleConfiguration) error {
	var err error
	var cli cmd.CLI
	var prv provisioner.Provisioner
	var ctx context.Context = context.Background()

	if configuration.Provisioner.Name == helm {
		cli = cmd.CLI{Binary: "helm"}
		prv = &provisioner.Helm{}
	}

	err = prv.Converge(ctx, &cli)

	return err
}
