package cmd

import (
	"strings"

	"github.com/apex/log"
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
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")

	logger := helpers.GetLogger(debug)

	logger.Info("Begin converge")

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Converge(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Converge(scenario, configuration, logger)

	logger.Info("Converge finished")
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
