package cmd

import (
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func verify(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	err = cli.Verify(scenario, configuration)
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
