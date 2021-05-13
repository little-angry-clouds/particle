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

func cleanup(cmd *cobra.Command, args []string) {
	var scenario string
	var configuration config.ParticleConfiguration
	var err error
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)
	scenario, _ = cmd.Flags().GetString("scenario")

	logger.Info("Begin cleanup")

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Cleanup(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Cleanup finished")
}

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "A brief description of your command",
	Run:   cleanup,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
