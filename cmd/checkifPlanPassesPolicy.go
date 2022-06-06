package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/tfval/model"
	opa "github.com/bradmccoydev/tfval/pkg/opa"

	"github.com/spf13/cobra"
)

var (
	checkifPlanPassesPolicyCmd = &cobra.Command{
		Use:   "check",
		Short: "check If plan passes policy",
		Long:  `Check if the plan passes Policy`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkifPlanPassesPolicy(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkifPlanPassesPolicyCmd)
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&opaConfig, "opaConfig", "o", opaConfig, "OPA Config")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
}

func checkifPlanPassesPolicy(args []string) string {
	plan, err := ioutil.ReadFile(planFileName)
	if err != nil {
		fmt.Println(err)
	}

	b := []byte(opaConfig)
	var config model.OpaConfig

	if err := json.Unmarshal(b, &config); err != nil {
		fmt.Println(err)
	}

	for _, policy := range config {
		policyResponse := opa.RetrieveOpaPolicyResponse(plan, policy.Location, policy.Query)

		b := []byte(policyResponse)
		var validations model.TfValidation

		if err := json.Unmarshal(b, &validations); err != nil {
			fmt.Println(err)
		}

		weights := opa.GetWeights(policyResponse)
		response := fmt.Sprintf("%d", validations[0].Score)

		for _, validation := range validations {
			score := opa.GetWeightByServiceNameAndAction(weights, validation.Data.Type, validation.Data.Change.Actions[0])
			response = response + fmt.Sprintf("%s %t %d ", validation.Data.Address, validation.ValidationPassed, score)

			//fmt.Println()
		}

		fmt.Println(response)

	}

	return ""
}
