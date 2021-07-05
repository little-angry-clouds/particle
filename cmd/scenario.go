package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func scenario(cmd *cobra.Command, args []string) {
	var scenario string = args[0]
	var driver string
	var provisioner string
	var lint string = "set -e\nhelm lint"
	var verifier string = "set -e\nhelm test "
	var debug bool
	var err error
	var configuration config.ParticleConfiguration

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)

	scenarioPath := "particle/" + scenario

	logger.Info("Begin initialization")

	driver, err = cmd.Flags().GetString("driver")
	customError.CheckGenericError(logger, err)

	provisioner, err = cmd.Flags().GetString("provisioner")
	customError.CheckGenericError(logger, err)

	path, err := os.Getwd()
	customError.CheckGenericError(logger, err)

	chartName := filepath.Base(path)
	verifier += chartName

	// Check if the particle directory exists and exit if not
	_, err = os.Stat("particle/")
	if os.IsNotExist(err) {
		_, err = os.Stat("particle/")
		if !os.IsNotExist(err) {
			customError.CheckGenericError(logger, err)
		}
	}

	// Check if the scenario directory exists and create it if not
	_, err = os.Stat(scenarioPath)
	if !os.IsNotExist(err) {
		customError.CheckGenericError(logger, fmt.Errorf("the scenario directory is already created, exiting"))
	}

	err = os.MkdirAll(scenarioPath, 0755)
	customError.CheckGenericError(logger, err)

	// Create particle configuration
	configuration.Driver.Name = driver
	configuration.Provisioner.Name = provisioner
	configuration.Linter = lint
	configuration.Verifier = verifier
	configuration.Dependency.Name = helm

	logger.WithFields(log.Fields{
		"driver":      driver,
		"provisioner": helm,
		"verifier":    strings.Replace(configuration.Verifier, "\n", " && ", -1),
		"linter":      strings.Replace(lint, "\n", " && ", -1),
	}).Debug("Configuration to create")

	err = config.CreateConfiguration("./"+scenarioPath, configuration)
	customError.CheckGenericError(logger, err)

	logger.Info("Initialization finished")
}

// scenarioCmd represents the scenario command
var scenarioCmd = &cobra.Command{
	Use:   "scenario",
	Short: "Initializes an scenario",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing argument 'SCENARIO_NAME'")
		}
		return nil
	},
	Run: scenario,
}

func init() {
	scenarioCmd.PersistentFlags().StringP("driver", "d", "kind", "driver to use when creating the kubernetes cluster")
	scenarioCmd.PersistentFlags().StringP("provisioner", "p", "helm", "provisioner to use when deploying to the kubernetes cluster")
	initCmd.AddCommand(scenarioCmd)
}
