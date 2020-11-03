package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// idempotenceCmd represents the idempotence command
var idempotenceCmd = &cobra.Command{
	Use:   "idempotence",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("idempotence called")
	},
}

func init() {
	rootCmd.AddCommand(idempotenceCmd)
}
