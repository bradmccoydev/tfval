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
	checkifPlanPassesOpaPolicyCmd = &cobra.Command{
		Use:   "checkopa",
		Short: "check If plan passes policy",
		Long:  `Check if the plan passes Policy`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkifPlanPassesOpaPolicy(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkifPlanPassesOpaPolicyCmd)
	checkifPlanPassesOpaPolicyCmd.PersistentFlags().StringVarP(&opaConfig, "opaConfig", "o", opaConfig, "OPA Config")
	checkifPlanPassesOpaPolicyCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
}

func checkifPlanPassesOpaPolicy(args []string) bool {
	plan, err := ioutil.ReadFile(planFileName)
	if err != nil {
		fmt.Println(err)
	}

	b := []byte(opaConfig)
	var o model.OpaConfig

	if err := json.Unmarshal(b, &o); err != nil {
		fmt.Println(err)
	}

	passesOpa := true

	for _, policy := range o {
		passesOpa = opa.CheckIfPlanPassesOpaPolicy(plan, policy.Location, policy.Query)
		if passesOpa == false {
			break
		}
	}

	return passesOpa
}
