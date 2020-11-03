package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sideEffectCmd represents the sideEffect command
var sideEffectCmd = &cobra.Command{
	Use:   "sideEffect",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sideEffect called")
	},
}

func init() {
	rootCmd.AddCommand(sideEffectCmd)
}
