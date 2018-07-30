package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nycdavid/kit/credentials"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCredCmd)
}

var setCredCmd = &cobra.Command{
	Use:   "setcred",
	Short: "Set credentials for kubectl",
	Long:  "Add a context to kubectl configured to the credentials in the current working directory's .kit.json file",
	Run: func(cmd *cobra.Command, args []string) {
		conf := credentials.ReadFile()

		kubectlArgs := []string{
			"config",
			"set-context",
			conf.Name,
			fmt.Sprintf("--cluster=%s", conf.Host),
			fmt.Sprintf("--user=%s", conf.Name),
			fmt.Sprintf("--namespace=default"),
		}
		runKubectlCmd(kubectlArgs)

		kubectlArgs = []string{"config", "use-context", conf.Name}
		runKubectlCmd(kubectlArgs)

		kubectlArgs = []string{
			"config",
			"set-credentials",
			conf.Name,
			fmt.Sprintf("--username=%s", conf.User),
			fmt.Sprintf("--password=%s", conf.Password),
		}
		fmt.Println(kubectlArgs)
		runKubectlCmd(kubectlArgs)

		kubectlArgs = []string{
			"config",
			"set-cluster",
			conf.Host,
			"--insecure-skip-tls-verify=true",
			fmt.Sprintf("--server=https://api.%s", conf.Host),
		}
		runKubectlCmd(kubectlArgs)
	},
}

func runKubectlCmd(kubectlArgs []string) {
	com := exec.Command("kubectl", kubectlArgs...)
	com.Stderr = os.Stderr
	com.Stdout = os.Stdout
	com.Run()
}
