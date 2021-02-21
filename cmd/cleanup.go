package cmd

import (
	"context"

	"github.com/spf13/cobra"

	c "github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

func cleanup(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var prv provisioner.Provisioner
	var ctx context.Context = context.Background()
	var cli c.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	if configuration.Provisioner.Name == helm {
		cli = c.CLI{Binary: "helm"}
		prv = &provisioner.Helm{}
	}

	err = prv.Cleanup(ctx, &cli)
	helpers.CheckGenericError(err)
}

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A brief description of your command",
	Run:   cleanup,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
