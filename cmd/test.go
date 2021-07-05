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

func test(cmd *cobra.Command, args []string) {
	var scenario string
	var err error
	var configuration config.ParticleConfiguration
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

	err = cli.Syntax(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Dependency")

	err = cli.Dependency(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Lint")

	err = cli.Lint(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Create")

	err = cli.Create(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Prepare")

	err = cli.Prepare(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Converge")

	err = cli.Converge(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Verify")

	err = cli.Verify(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Cleanup")

	err = cli.Cleanup(scenario, configuration, logger)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		logger.Warn("Destroy")
		_ = cli.Destroy(scenario, configuration, logger)

		return
	}

	logger.Info("Destroy")

	err = cli.Destroy(scenario, configuration, logger)
	customError.CheckGenericError(logger, err)
}

// cleanupCmd represents the cleanup command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Executes the full matrix of actions",
	Run:   test,
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.PersistentFlags().StringP("scenario", "s", "default", "scenario to use")
}
