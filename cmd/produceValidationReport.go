package cmd

import (
	"fmt"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var produceValidationReportCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce Validation Report",
	Long:  `Produce Terraform Validation Report`,
	Run: func(cmd *cobra.Command, args []string) {
		err := produceValidationReport(args)
		return err
		// if err != nil {
		// 	fmt.Println(err)
		// }
	},
}

var reportParams []string

func init() {
	rootCmd.AddCommand(produceValidationReportCmd)
	produceValidationReportCmd.PersistentFlags().StringArrayVar(&reportParams, "payload", reportParams, "slackwebhook")
}

func produceValidationReport(args []string) bool {
	fileName := args[0]
	cloudProvider := args[1]

	tfsecReport := tfsec.ProduceVulnerabilityReport(fileName)
	fmt.Println(tfsecReport)

	passesOpa := opa.CheckIfPlanPassesPolicy(fileName, cloudProvider)
	fmt.Println(passesOpa)

	if passesOpa {
		return true
	}

	return false
}
