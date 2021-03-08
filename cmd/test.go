package cmd

import (
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/cli"
	"github.com/little-angry-clouds/particle/internal/config"
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

	logger.Info("Begin test")

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(logger, err)

	logger.WithFields(log.Fields{
		"driver":      configuration.Driver.Name,
		"provisioner": configuration.Provisioner.Name,
		"verifier":    configuration.Verifier.Name,
		"lint":        strings.Replace(configuration.Lint, "\n", " && ", -1),
	}).Debug("Configuration to use")

	logger.Info("Begin lint")

	err = cli.Lint(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Lint(scenario, configuration, logger)

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Cleanup(scenario, configuration, logger)

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Destroy(scenario, configuration, logger)

	logger.Info("Create")

	err = cli.Create(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Create(scenario, configuration, logger)

	logger.Info("Converge")

	err = cli.Converge(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Converge(scenario, configuration, logger)

	logger.Info("Converge")

	err = cli.Verify(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Verify(scenario, configuration, logger)

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration)
	helpers.CheckGenericError(logger, err)
	err = cli.Cleanup(scenario, configuration, logger)

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration)
	helpers.CheckGenericError(logger, err)

	logger.Info("Test finished")
	err = cli.Destroy(scenario, configuration, logger)
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
