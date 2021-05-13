package cmd

import (
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func lint(cmd *cobra.Command, args []string) {
	var scenario string
	var err error
	var configuration config.ParticleConfiguration
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)
	scenario, _ = cmd.Flags().GetString("scenario")

	logger.Info("Begin linting")

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Lint(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

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
