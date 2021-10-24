package tfsec

import (
	"encoding/json"

	"github.com/bradmccoydev/terraform-plan-validator/model"
)

func ProduceVulnerabilityReport(plan []byte) model.Vulnerabilities {
	var vulnerabilities model.Vulnerabilities
	json.Unmarshal(plan, &vulnerabilities)

	return vulnerabilities
}

func CheckIfPlanPassesTfPolicy(plan []byte) bool {
	vulnerabilities := ProduceVulnerabilityReport(plan)
	passesPolicy := true

	if len(vulnerabilities.Results) > 0 {
		for _, element := range vulnerabilities.Results {
			if element.Severity == "CRITICAL" {
				passesPolicy = false
			}
		}
	}

	return passesPolicy
}
