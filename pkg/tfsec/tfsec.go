package tfsec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bradmccoydev/terraform-plan-validator/model"
)

func produceVulnerabilityReport(fileName string, cloudProvider string) string {
	report, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	var vulnerabilities model.Vulnerabilities
	json.Unmarshal([]byte(report), &vulnerabilities)

	if len(vulnerabilities.Results) > 0 {
		for _, element := range vulnerabilities.Results {
			fmt.Println(element.Severity)
			fmt.Println(cloudProvider)
		}
	}
	return "x"
}
