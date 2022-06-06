package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"

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

	passesOpa := ""

	for _, policy := range config {
		passesOpa = opa.RetrieveOpaPolicyResponse(plan, policy.Location, policy.Query)

	}

	return passesOpa
}
