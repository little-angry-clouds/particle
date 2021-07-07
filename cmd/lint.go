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

	err = cli.Syntax(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    strings.Replace(configuration.Verifier, "\n", " && ", -1),
		"linter":      strings.Replace(configuration.Linter, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Lint(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Linting finished")
}

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Executes arbitrary linters",
	Run:   lint,
}

func init() {
	rootCmd.AddCommand(lintCmd)
	lintCmd.PersistentFlags().StringP("scenario", "s", "default", "scenario to use")
}
