package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var produceValidationReportCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce Validation Report",
	Long:  `Produce Terraform Validation Report`,
	Run: func(cmd *cobra.Command, args []string) {
		err := produceValidationReport(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var reportParams []string

func init() {
	rootCmd.AddCommand(produceValidationReportCmd)
	produceValidationReportCmd.PersistentFlags().StringArrayVar(&reportParams, "payload", reportParams, "slackwebhook")
}

func produceValidationReport(args []string) error {
	fileName := args[2]
	cloudProvider := args[4]

	//var x = tfsec.produceVulnerabilityReport("x", "t")
	//fmt.Println(x)

	report, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var vulnerabilities model.Vulnerabilities
	json.Unmarshal([]byte(report), &vulnerabilities)

	if len(vulnerabilities.Results) > 0 {

		for _, element := range vulnerabilities.Results {
			fmt.Println(element.Severity)
			fmt.Println(cloudProvider)
		}
	}

	return nil
}

func produceVulnerabilityReport(fileName, cloudProvider string) {
	panic("unimplemented")
}
