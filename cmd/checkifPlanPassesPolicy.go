package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bradmccoydev/tfval/pkg/infracost"
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
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&infracostMonthlyBudget, "infracostMonthlyBudget", "b", infracostMonthlyBudget, "Monthly Budget")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&infracostReportLocation, "infracostReportLocation", "i", infracostReportLocation, "Infracost Report Location")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&repo, "repo", "e", repo, "Git Repository")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&commitSha, "commitSha", "m", commitSha, "Git Commit Sha")
	checkifPlanPassesPolicyCmd.PersistentFlags().StringVarP(&developer, "developer", "d", developer, "Developer")
}

func checkifPlanPassesPolicy(args []string) string {
	tfPlan := utils.ReadFile(planFileName)
	tfSecReport := utils.ReadFile(tfsecReportLocation)
	infracostReport := utils.ReadFile(infracostReportLocation)

	monthlyBudget, err := strconv.ParseFloat(infracostMonthlyBudget, 6)
	if err != nil {
		fmt.Println(err)
	}

	infracostResponse := infracost.CheckIfSpendIsWithinBudget([]byte(infracostReport), monthlyBudget)
	tfSecResponse := tfsec.CheckIfPlanPassesTfPolicy(tfSecReport, tfsecMaxSeverity)
	opaResponse := ""
	config := opa.GetOpaConfig(opaConfig)

	for _, policy := range config {
		policyResponse := opa.GetDefaultOpaResponse(tfPlan, policy.Location, policy.Query)
		opaResponse = fmt.Sprintf("%s%s,", opaResponse, policyResponse)
	}

	tfValPass := false
	if strings.Contains(tfSecResponse, "tfsec_pass\": true") && !strings.Contains(opaResponse, "opa_pass\":false") && strings.Contains(infracostResponse, "infracost_pass\":true") {
		tfValPass = true
	}

	opaResponse = strings.TrimRight(opaResponse, ",")
	return fmt.Sprintf("{\"tfval_pass\":%t,\"repo\":\"%s\",\"commit_sha\":\"%s\",\"developer\":\"%s\",\"tfsec\":%s,\"opa\":%s,\"infracost\":%s}", tfValPass, repo, commitSha, developer, tfSecResponse, opaResponse, infracostResponse)
}
