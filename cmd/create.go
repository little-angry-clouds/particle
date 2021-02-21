package cmd

import (
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func create(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	err = cli.Create(scenario, configuration)
	helpers.CheckGenericError(err)
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Use the provisioner to start the instances",
	Run:   create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
