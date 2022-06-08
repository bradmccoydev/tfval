package cmd

import (
	"fmt"
	"strconv"

	"github.com/bradmccoydev/tfval/pkg/infracost"
	"github.com/bradmccoydev/tfval/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	checkSpendIsWithinBudgetCmd = &cobra.Command{
		Use:   "cost",
		Short: "check if spend is within Budget",
		Long:  `Check if spend is within Budget`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkSpendIsWithinBudget(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkSpendIsWithinBudgetCmd)
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&infracostMonthlyBudget, "infracostMonthlyBudget", "b", infracostMonthlyBudget, "Monthly Budget")
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&infracostReportLocation, "infracostReportLocation", "i", infracostReportLocation, "Infracost Report Location")
}

//infracost breakdown --format json --terraform-var-file /deployment/dev.tfvars --path .
func checkSpendIsWithinBudget(args []string) string {
	infracostReport := utils.ReadFile(infracostReportLocation)

	monthlyBudget, err := strconv.ParseFloat(infracostMonthlyBudget, 6)
	if err != nil {
		fmt.Println(err)
	}

	infracostResponse := infracost.CheckIfSpendIsWithinBudget([]byte(infracostReport), monthlyBudget)
	return infracostResponse
}
