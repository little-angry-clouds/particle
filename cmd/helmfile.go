package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

func helmfile(cmd *cobra.Command, args []string) {
	var driver string
	var debug bool
	var err error
	var scenario string = "default"
	var provisioner string = "helmfile"
	var helmfileName string = args[0]
	var lint string = "set -e\nhelmfile lint"
	var verifier string = "set -e\nhelmfile test"
	var configuration config.ParticleConfiguration

	debug, _ = cmd.Flags().GetBool("debug")
	logger := helpers.GetLogger(debug)

	logger.Info("Begin initialization")

	driver, err = cmd.Flags().GetString("driver")
	customError.CheckGenericError(logger, err)

	// Check if the chart's directory exists and create it if not
	_, err = os.Stat(helmfileName)
	if !os.IsNotExist(err) {
		customError.CheckGenericError(logger, fmt.Errorf("the helm repository '%s' is already added", helmfileName))
	}

	err = os.MkdirAll(helmfileName+"/particle/"+scenario, 0755)
	customError.CheckGenericError(logger, err)

	// Create helmfile
	helmfile := `---
repositories:
- name: nginx
  url: https://kubernetes.github.io/ingress-nginx
helmDefaults:
  wait: true
  waitForJobs: true
releases:
  - name: nginx
    chart: nginx/ingress-nginx
    values:
      - controller:
          service:
            type: ClusterIP
`

	err = ioutil.WriteFile(helmfileName+"/helmfile.yaml", []byte(helmfile), 0644)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		os.Exit(1)

		return
	}

	// Create the configuration
	configuration.Driver.Name = driver
	configuration.Provisioner.Name = provisioner
	configuration.Linter = lint
	configuration.Verifier = verifier
	configuration.Dependency.Name = provisioner

	logger.WithFields(log.Fields{
		"driver":      driver,
		"provisioner": "helm",
		"verifier":    "helm",
		"linter":      strings.Replace(lint, "\n", " && ", -1),
	}).Debug("Configuration to create")

	err = config.CreateConfiguration(helmfileName+"/particle/"+scenario, configuration)
	customError.CheckGenericError(logger, err)

	logger.Info("Initialization finished")
}

// chartCmd represents the chart command
var helmfileCmd = &cobra.Command{
	Use:   "helmfile",
	Short: "Initializes a helmfile file and includes a default particle directory",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing argument 'HELMFILE_NAME'")
		}
		return nil
	},
	Run: helmfile,
}

func init() {
	helmfileCmd.PersistentFlags().StringP("driver", "d", "kind", "driver to use when creating the kubernetes cluster")
	helmfileCmd.PersistentFlags().StringP("provisioner", "p", "helm", "provisioner to use when deploying to the kubernetes cluster")
	initCmd.AddCommand(helmfileCmd)
}
