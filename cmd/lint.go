package cmd

import (
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func lint(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var err error
	var configuration config.ParticleConfiguration
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")

	logger := helpers.GetLogger(debug)

	logger.Info("Begin linting")

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Lint(scenario, configuration)
	helpers.CheckGenericError(logger, err)

	logger.Info("Linting finished")
}

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint executes external linters",
	Run:   lint,
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
