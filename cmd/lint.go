package cmd

import (
	"github.com/spf13/cobra"

	c "github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func lint(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var cli c.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	cli = c.CLI{Binary: "bash"}
	cmdArgs := []string{"bash", "-c", configuration.Lint}

	err = cli.Initialize(cmdArgs)
	helpers.CheckGenericError(err)

	err = cli.Run()
	helpers.CheckGenericError(err)
}

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint command executes external linters",
	Run:   lint,
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
