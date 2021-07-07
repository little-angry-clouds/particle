package cmd

import (
	"strings"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/spf13/cobra"
)

func prepare(cmd *cobra.Command, args []string) {
	var scenario string
	var configuration config.ParticleConfiguration
	var err error
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)
	scenario, _ = cmd.Flags().GetString("scenario")

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    strings.Replace(configuration.Verifier, "\n", " && ", -1),
		"linter":      strings.Replace(configuration.Linter, "\n", " && ", -1),
	}).Debug("Configuration to use")

	logger.Info("Syntax")

	err = cli.Syntax(configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Begin prepare")

	err = cli.Prepare(configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Prepare finished")
}

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Installs the dependencies",
	Run:   prepare,
}

func init() {
	rootCmd.AddCommand(prepareCmd)
	prepareCmd.PersistentFlags().StringP("scenario", "s", "default", "scenario to use")
}
