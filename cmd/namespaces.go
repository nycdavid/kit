package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(namespacesCmd)
}

type config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List all namespaces in your Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		var cnf config
		conf, err := os.Open(".preprod.kit.json")
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(conf)
		err = dec.Decode(&cnf)
		if err != nil {
			log.Fatal(err)
		}
		// set credentials
		com := exec.Command(
			"kubectl",
			"config",
			"set-credentials",
			fmt.Sprintf("%s-basicauth", cnf.Host),
			fmt.Sprintf("--username=%s", cnf.User),
			fmt.Sprintf("--password=%s", cnf.Password),
		)
		com.Stdout = os.Stdout
		com.Stderr = os.Stderr
		err = com.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// set cluster
		com = exec.Command(
			"kubectl",
			"config",
			"set-cluster",
			cnf.Host,
			"--insecure-skip-tls-verify=true",
			fmt.Sprintf("--server=https://api.%s", cnf.Host),
		)
		com.Stdout = os.Stdout
		com.Stderr = os.Stderr
		err = com.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		//

		// set context
		com = exec.Command(
			"kubectl",
			"config",
			"set-context",
			"default-context",
			fmt.Sprintf("--user=%s-basicauth", cnf.Host),
			"--namespace=default",
			fmt.Sprintf("--cluster=%s", cnf.Host),
		)
		com.Stdout = os.Stdout
		com.Stderr = os.Stderr
		err = com.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		//

		// use context
		com = exec.Command(
			"kubectl",
			"config",
			"use-context",
			"default-context",
		)
		com.Stdout = os.Stdout
		com.Stderr = os.Stderr
		err = com.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		//

		cmdN := exec.Command("kubectl", "get", "namespaces")
		cmdN.Stdout = os.Stdout
		cmdN.Stderr = os.Stderr
		err = cmdN.Run()
		if err != nil {
			os.Exit(1)
		}
	},
	ValidArgs: []string{"foo", "bar", "baz"},
}
