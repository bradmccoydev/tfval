package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var (
	tfsecPlanFileName  string
	tfsecCloudProvider string

	outputtfsecCmd = &cobra.Command{
		Use:   "tfsec",
		Short: "get tfsec report",
		Long:  `Get TfSec vulnerability report`,
		Run: func(cmd *cobra.Command, args []string) {
			result := outputTfsec(args, *cfg)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputtfsecCmd)
	outputtfsecCmd.PersistentFlags().StringVarP(&tfsecPlanFileName, "planFileName", "p", tfsecPlanFileName, "Plan file Name")
	outputtfsecCmd.PersistentFlags().StringVarP(&tfsecCloudProvider, "cloudProvider", "c", tfsecCloudProvider, "Cloud Provider")
}

func outputTfsec(args []string, cfg util.Config) model.Vulnerabilities {
	plan, err := ioutil.ReadFile(tfsecPlanFileName)
	if err != nil {
		fmt.Println(err)
	}

	tfsec := tfsec.OutputTfsecReport(plan, cfg)
	println(fmt.Sprint(tfsec))
	return tfsec
}
