package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// prepareCmd represents the prepare command
var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prepare called")
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
