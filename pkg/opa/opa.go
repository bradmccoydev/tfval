package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/open-policy-agent/opa/rego"
)

func CheckIfPlanPassesOpaPolicy(plan []byte, policyLocation string, opaRegoQuery string) bool {
	r := rego.New(
		rego.Query(opaRegoQuery),
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

func GetOpaScore(plan []byte, policyLocation string) int {
	r := rego.New(
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
