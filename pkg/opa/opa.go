package opa

import (
	"context"
	"encoding/json"
	"fmt"

	config "github.com/bradmccoydev/terraform-plan-validator/util"
	"github.com/open-policy-agent/opa/rego"
)

func CheckIfPlanPassesOpaPolicy(plan []byte, cloudProvider string, cfg config.Config) bool {
	policy := cfg.OpaGcpPolicy
	if cloudProvider == "azure" {
		policy = cfg.OpaAzurePolicy
	}

	r := rego.New(
		rego.Query("data.terraform.analysis.authz"),
		rego.Load([]string{policy}, nil))

	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		fmt.Println(err)
	}

	var input interface{}

	if err := json.Unmarshal(plan, &input); err != nil {
		fmt.Println(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
	}

	return rs.Allowed()

}
