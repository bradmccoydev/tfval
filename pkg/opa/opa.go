package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/open-policy-agent/opa/rego"
)

type Weight struct {
	Service string `json:"service"`
	Create  int    `json:"create"`
	Delete  int    `json:"delete"`
	Modify  int    `json:"modify"`
}

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

func GetWeights(payload string) []Weight {
	s := GetStringInBetweenTwoString(payload, "weights\":[", "}]")
	lines := strings.Split(s, "}")

	var weights []Weight

	for _, line := range lines {
		formatted := fmt.Sprintf("{ \"service\": %s%s", strings.Replace(line, ": {", ", ", -1), " },")
		formatted = strings.Replace(formatted, ": , \"", ": \"", -1)

		if strings.Contains(formatted, "create") {
			var weight Weight
			byte := []byte(strings.TrimRight(formatted, ","))

			if err := json.Unmarshal(byte, &weight); err != nil {
				fmt.Println(err)
			}

			weights = append(weights, weight)
		}
	}

	return weights
}

func GetWeightByServiceName(weights []Weight, serviceName string) int {
	return 1
}

func GetStringInBetweenTwoString(str string, startS string, endS string) (result string) {
	s := strings.Index(str, startS)
	//fmt.Println(strconv.Itoa(s) + "**")
	if s == -1 {
		return result
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}
