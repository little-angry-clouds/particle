package cmd

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/driver"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/spf13/cobra"
)

func destroy(cmd *cobra.Command, args []string) {
	// TODO add support to manage scenarios
	var scenario string = "default"
	var configuration config.ParticleConfiguration
	var err error
	var drv driver.Driver
	var ctx context.Context = context.Background()
	var cli driver.CLI

	configuration, err = config.ReadConfiguration(scenario)
	helpers.CheckGenericError(err)

	if configuration.Driver.Name == "kind" {
		cli = driver.CLI{}
		drv = &driver.Kind{}
	}

	err = drv.Destroy(ctx, &cli)
	helpers.CheckGenericError(err)
}

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Use the provisioner to destroy the instances",
	Run:   destroy,
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
