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
	var driver string
	var debug bool
	var err error
	var scenario string
	var provisioner string = "helm"
	var chartName string = args[0]
	var lint string = "set -e\nhelm lint"
	var verifier string = "set -e\nhelm test " + chartName
	var configuration config.ParticleConfiguration

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)

	logger.Info("Begin initialization")

	driver, err = cmd.Flags().GetString("driver")
	customError.CheckGenericError(logger, err)

	scenario, _ = cmd.Flags().GetString("scenario")

	// Check if the chart's directory exists and create it if not
	_, err = os.Stat(chartName)
	if !os.IsNotExist(err) {
		customError.CheckGenericError(logger, fmt.Errorf("the helm repository '%s' is already added", chartName))
	}

	err = os.MkdirAll(chartName+"/particle/"+scenario, 0755)
	customError.CheckGenericError(logger, err)

	// Create chart
	_, err = chartutil.Create(chartName, "")
	customError.CheckGenericError(logger, err)

	configuration.Driver.Name = driver
	configuration.Provisioner.Name = provisioner
	configuration.Linter = lint
	configuration.Verifier = verifier
	configuration.Dependency.Name = "helm"

	logger.WithFields(log.Fields{
		"driver":      driver,
		"provisioner": "helm",
		"verifier":    "helm",
		"linter":      strings.Replace(lint, "\n", " && ", -1),
	}).Debug("Configuration to create")

	err = config.CreateConfiguration(chartName+"/particle/"+scenario, configuration)
	customError.CheckGenericError(logger, err)

	logger.Info("Initialization finished")
}

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Initializes a helm chart and includes default particle directory",
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
	chartCmd.PersistentFlags().StringP("provisioner", "p", "helm", "provisioner to use when deploying to the kubernetes cluster")
	initCmd.AddCommand(chartCmd)
}
