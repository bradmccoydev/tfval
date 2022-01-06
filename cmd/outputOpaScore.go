package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var (
	opaScorePlanFileName  string
	opaScoreCloudProvider string

	outputOpaScoreCmd = &cobra.Command{
		Use:   "opascore",
		Short: "get opa score",
		Long:  `Get the OPA score report`,
		Run: func(cmd *cobra.Command, args []string) {
			result := outputOpaScore(args, *cfg)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputOpaScoreCmd)
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaScorePlanFileName, "planFileName", "p", opaScorePlanFileName, "Plan file Name")
	outputOpaScoreCmd.PersistentFlags().StringVarP(&opaScoreCloudProvider, "cloudProvider", "c", opaScoreCloudProvider, "Cloud Provider")
}

func outputOpaScore(args []string, cfg util.Config) int {
	plan, err := ioutil.ReadFile(opaScorePlanFileName)
	if err != nil {
		fmt.Println(err)
	}

	opaScore := opa.GetOpaScore(plan, opaScoreCloudProvider, cfg)
	return opaScore
}
