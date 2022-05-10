package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var (
	opaScorePlanFileName string
	opaPolicyLocation    string

	outputOpaScoreCmd = &cobra.Command{
		Use:   "opascore",
		Short: "get opa score",
		Long:  `Gets the OPA score report`,
		Run: func(cmd *cobra.Command, args []string) {
			result := outputOpaScore(args, *cfg)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputOpaScoreCmd)
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaScorePlanFileName, "planFileName", "p", opaScorePlanFileName, "Plan file Name")
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaPolicyLocation, "policyLocation", "c", opaPolicyLocation, "Cloud Provider")
}

func outputOpaScore(args []string, cfg util.Config) int {
	plan, err := ioutil.ReadFile(opaScorePlanFileName)
	if err != nil {
		fmt.Println(err)
	}

	opaScore := opa.GetOpaScore(plan, opaPolicyLocation, cfg)
	return opaScore
}
