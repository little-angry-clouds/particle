package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/chartutil"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func chart(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var driver string = "kind"
	var chartName string = args[0]
	var supportedDrivers = []string{driver}
	var err error
	var configuration config.ParticleConfiguration
	var lint string = "set -e\nhelm lint"

	driver, err = cmd.Flags().GetString("driver")
	helpers.CheckGenericError(err)

	if !helpers.StringInSlice(supportedDrivers, driver) {
		fmt.Printf("\"%s\" is not a valid value for the flag \"driver\".\n", driver)
		os.Exit(1)
	}

	// Check if the chart's directory exists and create it if not
	_, err = os.Stat(chartName)
	if !os.IsNotExist(err) {
		fmt.Println("Chart already exists.")
		os.Exit(1)
	}

	err = os.MkdirAll(chartName, 0755)
	helpers.CheckGenericError(err)

	// Create chart
	if _, err = chartutil.Create(chartName, ""); err != nil {
		fmt.Println("Could not create chart: ")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Chart created.")

	configuration.Driver.Name = driver
	configuration.Provisioner.Name = helm
	configuration.Lint = lint
	configuration.Verifier.Name = "helm"
	err = config.CreateConfiguration(chartName, scenario, configuration)
	helpers.CheckGenericError(err)
	fmt.Println("Particle initialized.")
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
