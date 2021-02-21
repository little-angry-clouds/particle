package cmd

import (
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func destroy(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	err = cli.Destroy(scenario, configuration)
	helpers.CheckGenericError(err)
}

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Use the provisioner to destroy the instances",
	Run:   destroy,
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
