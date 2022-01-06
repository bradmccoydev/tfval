package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	config "github.com/bradmccoydev/terraform-plan-validator/util"
	"github.com/open-policy-agent/opa/rego"
)

func CheckIfPlanPassesOpaPolicy(plan []byte, cloudProvider string, cfg config.Config) bool {
	policyLocation := cfg.OpaAzurePolicy
	if cloudProvider == "gcp" {
		policyLocation = cfg.OpaGcpPolicy
	}

	r := rego.New(
		rego.Query(cfg.OpaRegoQuery),
		rego.Load([]string{policyLocation}, nil))

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

func GetOpaScore(plan []byte, cloudProvider string, cfg config.Config) int {
	policyLocation := cfg.OpaAzurePolicy
	if cloudProvider == "gcp" {
		policyLocation = cfg.OpaGcpPolicy
	}

	r := rego.New(
		//rego.Query(cfg.OpaRegoQuery),
		rego.Query("data.terraform.analysis.score"),
		rego.Load([]string{policyLocation}, nil))

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
	s := fmt.Sprint(rs[0].Expressions[0].Value)
	i, _ := strconv.Atoi(s)
	return i
}
