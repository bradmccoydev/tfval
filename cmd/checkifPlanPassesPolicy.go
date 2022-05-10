package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var (
	planFileName   string
	policyLocation string

	checkifPlanPassesPolicyCmd = &cobra.Command{
		Use:   "check",
		Short: "check If plan passes policy",
		Long:  `Check if the plan passes Policy`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkifPlanPassesPolicy(args, *cfg)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkifPlanPassesPolicyCmd)
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&policyLocation, "policyLocation", "c", policyLocation, "Policy Location")
}

func checkifPlanPassesPolicy(args []string, cfg util.Config) bool {
	plan, err := ioutil.ReadFile(planFileName)
	if err != nil {
		fmt.Println(err)
	}

	passesTfsec := tfsec.CheckIfPlanPassesTfPolicy(plan, cfg)
	passesOpa := opa.CheckIfPlanPassesOpaPolicy(plan, policyLocation, cfg)

	if passesOpa && passesTfsec {
		return true
	}

	return false
}
