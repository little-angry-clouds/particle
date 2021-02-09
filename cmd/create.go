package cmd

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/driver"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/spf13/cobra"
)

func create(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var drv driver.Driver
	var ctx context.Context = context.Background()
	var kubernetesVersion config.Key = "kubernetesVersion"
	var cli driver.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	if configuration.Driver.Name == "kind" {
		cli = driver.CLI{}
		drv = &driver.Kind{}
	}

	// Pass variables to context
	if configuration.Driver.KubernetesVersion != "" {
		ctx = context.WithValue(ctx, kubernetesVersion, configuration.Driver.KubernetesVersion)
	}
	err = drv.Create(ctx, &cli)
	helpers.CheckGenericError(err)
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Use the provisioner to start the instances",
	Run:   create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
