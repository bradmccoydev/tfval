package cmd

import (
	"fmt"

	"github.com/bradmccoydev/tfval/model"
	tfsec "github.com/bradmccoydev/tfval/pkg/tfsec"
	"github.com/bradmccoydev/tfval/pkg/utils"

	"github.com/spf13/cobra"
)

var (
	outputTfSecCmd = &cobra.Command{
		Use:   "tfsecreport",
		Short: "get tfsec report",
		Long:  `Outputs TfSec vulnerability report`,
		Run: func(cmd *cobra.Command, args []string) {
			result := outputTfsec(args)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(outputTfSecCmd)
	outputTfSecCmd.PersistentFlags().StringVarP(&planFileName, "planFileName", "p", planFileName, "Plan file Name")
}

func outputTfsec(args []string) model.Vulnerabilities {
	plan := utils.ReadFile(fileName)

	tfsec := tfsec.OutputTfsecReport(plan)
	println(fmt.Sprint(tfsec))
	return tfsec
}
