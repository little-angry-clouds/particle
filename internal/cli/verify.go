package cli

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/verifier"
)

func Verify(scenario string, configuration config.ParticleConfiguration) error {
	var err error
	var cli cmd.CLI
	var vrf verifier.Verifier
	var ctx context.Context = context.Background()

	if configuration.Verifier.Name == "helm" {
		cli = cmd.CLI{Binary: "helm"}
		vrf = &verifier.Helm{}
	}

	err = vrf.Verify(ctx, &cli)

	return err
}
