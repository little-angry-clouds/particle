package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// convergeCmd represents the converge command
var convergeCmd = &cobra.Command{
	Use:   "converge",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("converge called")
	},
}

func init() {
	rootCmd.AddCommand(convergeCmd)
}
