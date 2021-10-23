package opa

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/open-policy-agent/opa/rego"
)

func main(fileName string, cloudProvider string) {
	policy := "./opa-gcp-policy.rego"
	if cloudProvider == "Azure" {
		policy = "./opa-azure-policy.rego"
	}

	r := rego.New(
		rego.Query("data.terraform.analysis.authz"),
		rego.Load([]string{policy}, nil))

	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		fmt.Println(err)
	}

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		fmt.Println(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rs[0])
}
