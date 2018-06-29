package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(namespacesCmd)
}

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List all namespaces in your Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmdN := exec.Command("kubectl", "get", "namespaces")
		com.Stdout = os.Stdout
		com.Stderr = os.Stderr
		err := com.Run()
		if err != nil {
			os.Exit(1)
		}
	},
}
