package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a pod",
	Long:  "SSH into a running pod on your cluster",
	Run: func(cmd *cobra.Command, args []string) {
		com := exec.Command("kubectl", "get", "namespaces")
		com.Stderr = os.Stderr
		com.Stdout = os.Stdout
		com.Run()
	},
}
