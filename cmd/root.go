package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "particle",
	Short: "Particle is a project designed to aid in the development and testing of kubernetes resources",
}

func Execute() {
	rootCmd.PersistentFlags().BoolP("debug", "D", false, "when true gets more verbose logs")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
