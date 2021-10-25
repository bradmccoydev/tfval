package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var checkifPlanPassesPolicyCmd = &cobra.Command{
	Use:   "check",
	Short: "check If plan passes policy",
	Long:  `Check if plan passes Policy`,
	Run: func(cmd *cobra.Command, args []string) {
		result := checkifPlanPassesPolicy(args, *cfg)
		fmt.Println(result)
	},
}

var planParams []string

func init() {
	rootCmd.AddCommand(checkifPlanPassesPolicyCmd)
	checkifPlanPassesPolicyCmd.PersistentFlags().StringArrayVar(&planParams, "payload", planParams, "slackwebhook")
}

func checkifPlanPassesPolicy(args []string, cfg util.Config) bool {
	fileName := args[0]
	cloudProvider := args[1]

	plan, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	passesTfsec := tfsec.CheckIfPlanPassesTfPolicy(plan)
	fmt.Println("passes:", passesTfsec)

	passesOpa := opa.CheckIfPlanPassesOpaPolicy(plan, cloudProvider, cfg)
	fmt.Println("passes: ", passesOpa)

	if passesOpa && passesTfsec {
		return true
	}

	return false
}
