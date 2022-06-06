package tfsec

import (
	"encoding/json"
	"fmt"

	"github.com/bradmccoydev/tfval/model"
)

func ProduceVulnerabilityReport(plan []byte) model.Vulnerabilities {
	var vulnerabilities model.Vulnerabilities
	json.Unmarshal(plan, &vulnerabilities)

	return vulnerabilities
}

func CheckIfPlanPassesTfPolicy(plan []byte, tfsecMaxSeverity string) bool {
	vulnerabilities := ProduceVulnerabilityReport(plan)
	passesPolicy := true

	if IsInvalidCategory(tfsecMaxSeverity) {
		fmt.Println("Invalid TFSEC Severity Category defaulting to LOW")
		tfsecMaxSeverity = "LOW"
	}

	if tfsecMaxSeverity == "LOW" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "LOW" || element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
				}
			}
		}
	}

	if tfsecMaxSeverity == "MEDIUM" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
				}
			}
		}
	}

	if tfsecMaxSeverity == "CRITICAL" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "CRITICAL" {
					passesPolicy = false
				}
			}
		}
	}

	return passesPolicy
}

func OutputTfsecReport(tfsecJsonOutput []byte) model.Vulnerabilities {
	vulnerabilities := ProduceVulnerabilityReport(tfsecJsonOutput)

	return vulnerabilities
}

func IsInvalidCategory(category string) bool {
	switch category {
	case
		"LOW",
		"MEDIUM",
		"CRITICAL":
		return false
	}
	return true
}
