package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var checkifPlanPassesPolicyCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce Validation Report",
	Long:  `Produce Terraform Validation Report`,
	Run: func(cmd *cobra.Command, args []string) {
		result := checkifPlanPassesPolicy(args)
		fmt.Println(result)
	},
}

var reportParams []string

func init() {
	rootCmd.AddCommand(checkifPlanPassesPolicyCmd)
	checkifPlanPassesPolicyCmd.PersistentFlags().StringArrayVar(&reportParams, "payload", reportParams, "slackwebhook")
}

func checkifPlanPassesPolicy(args []string) bool {
	fileName := args[0]
	cloudProvider := args[1]

	plan, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	passesTfsec := tfsec.CheckIfPlanPassesTfPolicy(plan)
	fmt.Println("passes:", passesTfsec)

	passesOpa := opa.CheckIfPlanPassesOpaPolicy(plan, cloudProvider)
	fmt.Println("passes: ", passesOpa)

	if passesOpa && passesTfsec {
		return true
	}

	return false
}
