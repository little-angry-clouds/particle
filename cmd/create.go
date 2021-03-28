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

func create(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")

	logger := helpers.GetLogger(debug)

	logger.Info("Begin create")

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err, true)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Create(scenario, configuration, logger)
	customError.CheckGenericError(logger, err, true)

	logger.Info("Create finished")
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Use the provisioner to start the instances",
	Run:   create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
