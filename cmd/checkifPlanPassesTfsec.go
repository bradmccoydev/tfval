package cmd

import (
	"fmt"
	"io/ioutil"

	tfsec "github.com/bradmccoydev/tfval/pkg/tfsec"

	"github.com/spf13/cobra"
)

var (
	checkifPlanPassesTfSecCmd = &cobra.Command{
		Use:   "tfsec",
		Short: "check If passes tfsec policy",
		Long:  `Check if passes tfsec Policy`,
		Run: func(cmd *cobra.Command, args []string) {
			result := checkifPlanPassesTfSec(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(checkifPlanPassesTfSecCmd)
	checkifPlanPassesTfSecCmd.PersistentFlags().StringVarP(&tfsecReportLocation, "tfsecReportLocation", "t", tfsecReportLocation, "Tfsec Report")
	checkifPlanPassesTfSecCmd.PersistentFlags().StringVarP(&tfsecMaxSeverity, "tfsecMaxSeverity", "s", tfsecMaxSeverity, "The TF Sec Max Severity")
}

func checkifPlanPassesTfSec(args []string) string {
	report, err := ioutil.ReadFile(tfsecReportLocation)
	if err != nil {
		fmt.Println(err)
	}

	return tfsec.CheckIfPlanPassesTfPolicy(report, tfsecMaxSeverity)
}
