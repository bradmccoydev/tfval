package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/bradmccoydev/tfval/model"
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
	checkSpendIsWithinBudgetCmd.PersistentFlags().StringVarP(&opaConfig, "opaConfig", "p", opaConfig, "OPA Config")
}

//infracost breakdown --format json --terraform-var-file /deployment/dev.tfvars --path .
func checkSpendIsWithinBudget(args []string) bool {
	b := []byte(opaConfig)
	var o model.OpaConfig

	if err := json.Unmarshal(b, &o); err != nil {
		fmt.Println(err)
	}

	for _, element := range o {
		fmt.Println(element.Location)
		fmt.Println(element.Query)
	}

	return false
}
