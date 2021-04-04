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

func syntax(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")

	logger := helpers.GetLogger(debug)

	logger.Info("Begin syntax")

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	err = cli.Syntax(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Syntax finished")
}

// syntaxCmd represents the syntax command
var syntaxCmd = &cobra.Command{
	Use:   "syntax",
	Short: "Checks that the particle configuration is correct.",
	Run:   syntax,
}

func init() {
	rootCmd.AddCommand(syntaxCmd)
}
