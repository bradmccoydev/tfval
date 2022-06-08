package cmd

import (
	"fmt"
	"strings"

	opa "github.com/bradmccoydev/tfval/pkg/opa"
	tfsec "github.com/bradmccoydev/tfval/pkg/tfsec"
	"github.com/bradmccoydev/tfval/pkg/utils"

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
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&tfsecReportLocation, "tfsecReportLocation", "r", tfsecReportLocation, "Tfsec report location")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&tfsecMaxSeverity, "tfsecMaxSeverity", "s", tfsecMaxSeverity, "Tfsec Max Severity")
}

func checkifPlanPassesPolicy(args []string) string {
	tfReport := utils.ReadFile(tfsecReportLocation)
	tfPlan := utils.ReadFile(planFileName)

	tfSecResponse := tfsec.CheckIfPlanPassesTfPolicy(tfReport, tfsecMaxSeverity)
	opaResponse := ""

	config := opa.GetOpaConfig(opaConfig)

	for _, policy := range config {
		policyResponse := opa.GetDefaultOpaResponse(tfPlan, policy.Location, policy.Query)
		opaResponse = fmt.Sprintf("%s%s,", opaResponse, policyResponse)
	}

	tfValPass := false
	if strings.Contains(tfSecResponse, "tfsec_pass\": true") && strings.Contains(opaResponse, "opa_validation_passed\":false") {
		tfValPass = true
	}

	opaResponse = strings.TrimRight(opaResponse, ",")
	return fmt.Sprintf("{\"tfval_pass\":%t,\"tfsec\":%s,\"opa\":%s}", tfValPass, tfSecResponse, opaResponse)
}
