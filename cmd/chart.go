package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/chartutil"

	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func chart(cmd *cobra.Command, args []string) {
	var scenario string
	var driver string
	var chartName string = args[0]
	var lint string = "set -e\nhelm lint"
	var debug bool
	var err error
	var configuration config.ParticleConfiguration

	debug, _ = cmd.Flags().GetBool("debug")
	scenario, _ = cmd.Flags().GetString("scenario")
	logger := helpers.GetLogger(debug)

	logger.Info("Begin initialization")

	driver, err = cmd.Flags().GetString("driver")
	customError.CheckGenericError(logger, err)

	// Check if the chart's directory exists and create it if not
	_, err = os.Stat(chartName)
	if !os.IsNotExist(err) {
		customError.CheckGenericError(logger, fmt.Errorf("the helm repository '%s' is already added", chartName))
	}

	err = os.MkdirAll(chartName, 0755)
	customError.CheckGenericError(logger, err)

	// Create chart
	_, err = chartutil.Create(chartName, "")
	customError.CheckGenericError(logger, err)

	configuration.Driver.Name = driver
	configuration.Provisioner.Name = helm
	configuration.Lint = lint
	configuration.Verifier.Name = helm
	configuration.Dependency.Name = helm

	logger.WithFields(log.Fields{
		"driver":      driver,
		"provisioner": helm,
		"verifier":    helm,
		"lint":        strings.Replace(lint, "\n", " && ", -1),
	}).Debug("Configuration to create")

	err = config.CreateConfiguration(chartName, scenario, configuration)
	customError.CheckGenericError(logger, err)

	logger.Info("Initialization finished")
}

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Initialize a helm chart and include default particle directory.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing argument 'CHART_NAME'")
		}
		return nil
	},
	Run: chart,
}

func init() {
	chartCmd.PersistentFlags().StringP("driver", "d", "kind", "driver to use when creating the kubernetes cluster")
	initCmd.AddCommand(chartCmd)
}
