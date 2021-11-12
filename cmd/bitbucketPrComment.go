package cmd

import (
	"fmt"
	"io/ioutil"

	opa "github.com/bradmccoydev/terraform-plan-validator/pkg/opa"
	tfsec "github.com/bradmccoydev/terraform-plan-validator/pkg/tfsec"
	"github.com/bradmccoydev/terraform-plan-validator/util"

	"github.com/spf13/cobra"
)

var (
	tfPlanFileName  string
	tfCloudProvider string

	bitbucketPrCommentCmd = &cobra.Command{
		Use:   "check",
		Short: "check If plan passes policy and comment",
		Long:  `Check if plan passes Policy and comment on Pull Request`,
		Run: func(cmd *cobra.Command, args []string) {
			result := bitbucketPrComment(args, *cfg)
			fmt.Println(result)
		},
	}
)

func init() {
	rootCmd.AddCommand(bitbucketPrCommentCmd)
	bitbucketPrCommentCmd.PersistentFlags().StringVarP(&tfPlanFileName, "PlanFileName", "p", tfPlanFileName, "Plan file Name")
	bitbucketPrCommentCmd.PersistentFlags().StringVarP(&tfCloudProvider, "CloudProvider", "c", tfCloudProvider, "Cloud Provider")
}

func bitbucketPrComment(args []string, cfg util.Config) bool {
	plan, err := ioutil.ReadFile(tfPlanFileName)
	if err != nil {
		fmt.Println(err)
	}

	passesTfsec := tfsec.CheckIfPlanPassesTfPolicy(plan, cfg)
	passesOpa := opa.CheckIfPlanPassesOpaPolicy(plan, tfCloudProvider, cfg)

	if passesOpa && passesTfsec {
		return true
	}

	return false
}
