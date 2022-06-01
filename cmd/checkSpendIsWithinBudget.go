package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"

	"github.com/spf13/cobra"
)

var (
	planFileName   string
	policyLocation string
	opaRegoQuery   string

	checkSpendIsWithinBudgetCmd = &cobra.Command{
		Use:   "check",
		Short: "check If plan passes policy",
		Long:  `Check if the plan passes Policy`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkSpendIsWithinBudget(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkSpendIsWithinBudgetCmd)
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&policyLocation, "policyLocation", "c", policyLocation, "Policy Location")
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&tfsecMaxSeverity, "tfsecMaxSeverity", "s", tfsecMaxSeverity, "The TF Sec Max Severity")
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&opaRegoQuery, "opaRegoQuery", "o", opaRegoQuery, "The TF Sec Max Severity")
	//checkSpendIsWithinBudgetCmd.PersistentFlags().StringArrayVarP()
}

func checkSpendIsWithinBudget(args []string) bool {
	plan, err := ioutil.ReadFile(planFileName)
	if err != nil {
		fmt.Println(err)
	}

	passesTfsec := tfsec.CheckIfPlanPassesTfPolicy(plan, tfsecMaxSeverity)
	passesOpa := opa.CheckIfPlanPassesOpaPolicy(plan, policyLocation, opaRegoQuery)

	if passesOpa && passesTfsec {
		return true
	}

	return false
}
