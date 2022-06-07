package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bradmccoydev/tfval/model"
	"github.com/open-policy-agent/opa/rego"
)

func RetrieveOpaPolicyResponse(plan []byte, policyLocation string, opaRegoQuery string) string {
	rs := GetOpaResultSet(plan, policyLocation, opaRegoQuery)
	response := fmt.Sprintf("%v", rs[0].Expressions)

	return strings.Replace(response, "},]", "}]", -1)
}

func CheckIfPlanPassesOpaPolicy(plan []byte, policyLocation string, opaRegoQuery string) bool {
	rs := GetOpaResultSet(plan, policyLocation, opaRegoQuery)

	// fmt.Println(rs[0].Expressions)
	// fmt.Println(rs.Allowed())

	return rs.Allowed()
}

func GetOpaScore(plan []byte, policyLocation string, opaRegoQuery string) int {
	rs := GetOpaResultSet(plan, policyLocation, opaRegoQuery)

	s := fmt.Sprint(rs[0].Expressions[0].Value)
	i, _ := strconv.Atoi(s)
	return i
}

func GetOpaResultSet(plan []byte, policyLocation string, opaRegoQuery string) rego.ResultSet {
	ctx := context.Background()
	var input interface{}

	r := rego.New(
		rego.Query(opaRegoQuery),
		rego.Load([]string{policyLocation}, nil))

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(plan, &input); err != nil {
		fmt.Println(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
	}

	return rs
}

func GetTfWeights(payload string) []model.Weight {
	s := GetStringInBetweenTwoString(payload, "weights\":[", "}]")
	lines := strings.Split(s, "}")

	var weights []model.Weight

	for _, line := range lines {
		formatted := fmt.Sprintf("{ \"service\": %s%s", strings.Replace(line, ": {", ", ", -1), " },")
		formatted = strings.Replace(formatted, ": , \"", ": \"", -1)

		if strings.Contains(formatted, "create") {
			var weight model.Weight
			byte := []byte(strings.TrimRight(formatted, ","))

			if err := json.Unmarshal(byte, &weight); err != nil {
				fmt.Println(err)
			}

			weights = append(weights, weight)
		}
	}

	return weights
}

func GetTfWeightByServiceNameAndAction(weights []model.Weight, serviceName string, action string) int {
	actionWeight := 0

	for _, weight := range weights {
		if weight.Service == serviceName {
			switch action {
			case "delete":
				actionWeight = weight.Delete
			case "create":
				actionWeight = weight.Create
			case "modify":
				actionWeight = weight.Modify
			default:
				actionWeight = 0
			}
		}
	}
	return actionWeight
}

func GetStringInBetweenTwoString(str string, startS string, endS string) (result string) {
	s := strings.Index(str, startS)
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
