package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dependency called")
	},
}

func init() {
	rootCmd.AddCommand(dependencyCmd)
}
