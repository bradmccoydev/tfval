package cmd

import (
	"fmt"
	"os"

	"github.com/bradmccoydev/terraform-plan-validator/util"
	config "github.com/bradmccoydev/terraform-plan-validator/util"
	"github.com/spf13/cobra"
)

var cfgFile string
var cfg *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "Terraform Plan Validator",
	Short:   "Terraform Plan Validator",
	Long:    `Terraform Plan Validator`,
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	conf, err := util.LoadConfig(".")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	cfg = &config.Config{}
	cfg.OpaGcpPolicy = conf.OpaGcpPolicy
	cfg.OpaAzurePolicy = conf.OpaAzurePolicy
	cfg.OpaAwsPolicy = conf.OpaAwsPolicy
}
