package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chart called")
	},
}

func init() {
	initCmd.AddCommand(chartCmd)
}
