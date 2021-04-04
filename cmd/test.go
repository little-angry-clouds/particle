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

func test(cmd *cobra.Command, args []string) { // nolint: funlen
	// TODO add support to manage scenarios
	var scenario string = "default"
	var err error
	var configuration config.ParticleConfiguration
	var debug bool

	debug, _ = cmd.Flags().GetBool("debug")

	logger := helpers.GetLogger(debug)

	configuration, err = config.ReadConfiguration(scenario)
	customError.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	logger.Info("Syntax")

	err = cli.Syntax(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Dependency")

	err = cli.Dependency(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Lint")

	err = cli.Lint(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Create")

	err = cli.Create(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Prepare")

	err = cli.Prepare(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Converge")

	err = cli.Converge(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Verify")

	err = cli.Verify(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)
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
