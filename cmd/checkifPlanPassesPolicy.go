package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/tfval/model"
	opa "github.com/bradmccoydev/tfval/pkg/opa"
	tfsec "github.com/bradmccoydev/tfval/pkg/tfsec"

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
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&opaConfig, "opaConfig", "o", opaConfig, "OPA Config")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&tfsecReportLocation, "tfsecReportLocation", "p", tfsecReportLocation, "Tfsec report location")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&tfsecMaxSeverity, "tfsecMaxSeverity", "p", tfsecMaxSeverity, "Tfsec Max Severity")
}

func checkifPlanPassesPolicy(args []string) string {
	report, err := ioutil.ReadFile(tfsecReportLocation)
	if err != nil {
		fmt.Println(err)
		return "{\"error\": \"Error reading TfSec Report\"}"
	}

	tfResponse := tfsec.CheckIfPlanPassesTfPolicy(report, tfsecMaxSeverity)

	tfPlan, err := ioutil.ReadFile(planFileName)
	if err != nil {
		fmt.Println(err)
		return "{\"error\": \"Error reading Tfplan\"}"
	}

	b := []byte(opaConfig)
	var config model.OpaConfig

	if err := json.Unmarshal(b, &config); err != nil {
		fmt.Println(err)
		return "{\"error\": \"Error reading OpaConfig\"}"
	}

	opaResponse := ""

	for _, policy := range config {
		policyResponse := opa.RetrieveOpaPolicyResponse(tfPlan, policy.Location, policy.Query)

		b := []byte(policyResponse)
		var validations model.TfValidation

		if err := json.Unmarshal(b, &validations); err != nil {
			fmt.Println(err)
			return "{\"error\": \"Error reading Tf Validations\"}"
		}

		weights := opa.GetTfWeights(policyResponse)
		opaResponse = fmt.Sprintf("%d", validations[0].Score)

		for _, validation := range validations {
			score := opa.GetTfWeightByServiceNameAndAction(weights, validation.Data.Type, validation.Data.Change.Actions[0])
			opaResponse = fmt.Sprintf("%s %s %t %d ", opaResponse, validation.Data.Address, validation.ValidationPassed, score)
		}

		fmt.Println(tfResponse)

	}

	return fmt.Sprintf("{\"tfsec\": \"%s\", ", tfResponse)
}
