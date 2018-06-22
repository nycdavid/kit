package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "KIT is a collection of tools for a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("you ran kit")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
