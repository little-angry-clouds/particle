package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lint called")
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
