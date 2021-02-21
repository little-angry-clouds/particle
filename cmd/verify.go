package cmd

import (
	"context"

	"github.com/spf13/cobra"

	c "github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/little-angry-clouds/particle/internal/verifier"
)

func verify(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var vrf verifier.Verifier
	var ctx context.Context = context.Background()
	var cli c.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	if configuration.Verifier.Name == "helm" {
		cli = c.CLI{Binary: "helm"}
		vrf = &verifier.Helm{}
	}

	err = vrf.Verify(ctx, &cli)
	helpers.CheckGenericError(err)
}

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify that the state is what we want",
	Run:   verify,
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
