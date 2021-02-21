package cmd

import (
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func test(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var err error
	var configuration config.ParticleConfiguration

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	err = cli.Lint(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Cleanup(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Destroy(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Create(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Converge(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Verify(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Cleanup(scenario, configuration)
	helpers.CheckGenericError(err)

	err = cli.Destroy(scenario, configuration)
	helpers.CheckGenericError(err)
}

// cleanupCmd represents the cleanup command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute the full matrix of actions",
	Run:   test,
}

func init() {
	rootCmd.AddCommand(testCmd)
}
