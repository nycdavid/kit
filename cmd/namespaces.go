package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(namespacesCmd)
}

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List all namespaces in your Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is kit namespaces")
	},
}
