package cmd

import (
	"fmt"

	opa "github.com/bradmccoydev/tfval/pkg/opa"
	"github.com/bradmccoydev/tfval/pkg/utils"

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
			result := outputOpaScore(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputOpaScoreCmd)
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaScorePlanFileName, "planFileName", "p", opaScorePlanFileName, "Plan file Name")
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaPolicyLocation, "policyLocation", "c", opaPolicyLocation, "Cloud Provider")
}

func outputOpaScore(args []string) int {
	plan := utils.ReadFile(opaScorePlanFileName)

	opaScore := opa.GetOpaScore(plan, opaPolicyLocation, "data.terraform.analysis.score")
	return opaScore
}
