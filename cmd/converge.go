package cmd

import (
	"context"

	"github.com/spf13/cobra"

	c "github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/little-angry-clouds/particle/internal/provisioner"
)

func converge(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var prv provisioner.Provisioner
	var ctx context.Context = context.Background()
	var cli c.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	if configuration.Provisioner.Name == "helm" {
		cli = c.CLI{Binary: "helm"}
		prv = &provisioner.Helm{}
	}

	err = prv.Converge(ctx, &cli)
	helpers.CheckGenericError(err)
}

// convergeCmd represents the converge command
var convergeCmd = &cobra.Command{
	Use:   "converge",
	Short: "Converge will execute the sequence necessary to converge the instances",
	Run:   converge,
}

func init() {
	rootCmd.AddCommand(convergeCmd)
}
