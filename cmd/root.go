package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	tfsecMaxSeverity string
	planFileName     string

	rootCmd = &cobra.Command{
		Use:     "Terraform Plan Validator",
		Short:   "Terraform Plan Validator",
		Long:    `Terraform Plan Validator`,
		Version: "0.0.2",
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
