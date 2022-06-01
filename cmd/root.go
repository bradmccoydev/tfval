package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	tfsecReportLocation string
	tfsecMaxSeverity    string
	planFileName        string
	opaConfig           string

	rootCmd = &cobra.Command{
		Use:     "tfval",
		Short:   "Terraform Validator",
		Long:    `Terraform Validator`,
		Version: "1.1.2",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
