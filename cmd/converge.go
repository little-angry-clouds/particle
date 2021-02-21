package cmd

import (
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func converge(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	err = cli.Converge(scenario, configuration)
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
