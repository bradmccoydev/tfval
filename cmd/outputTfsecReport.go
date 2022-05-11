package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"

	"github.com/spf13/cobra"
)

var (
	tfsecPlanFileName string

	outputtfsecCmd = &cobra.Command{
		Use:   "tfsec",
		Short: "get tfsec report",
		Long:  `Outputs TfSec vulnerability report`,
		Run: func(cmd *cobra.Command, args []string) {
			result := outputTfsec(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputtfsecCmd)
	outputtfsecCmd.PersistentFlags().StringVarP(&tfsecPlanFileName, "planFileName", "p", tfsecPlanFileName, "Plan file Name")
}

func outputTfsec(args []string) model.Vulnerabilities {
	plan, err := ioutil.ReadFile(tfsecPlanFileName)
	if err != nil {
		fmt.Println(err)
	}

	tfsec := tfsec.OutputTfsecReport(plan)
	println(fmt.Sprint(tfsec))
	return tfsec
}
