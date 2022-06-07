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

func CheckIfPlanPassesTfPolicy(plan []byte, tfsecMaxSeverity string) string {
	vulnerabilities := ProduceVulnerabilityReport(plan)
	passesPolicy := true

	if IsInvalidCategory(tfsecMaxSeverity) {
		fmt.Println("Invalid TFSEC Severity Category defaulting to LOW")
		tfsecMaxSeverity = "LOW"
	}

	vulnerabilitiesMessage := ""

	if tfsecMaxSeverity == "LOW" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "LOW" || element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
					vulnerabilitiesMessage = fmt.Sprintf("%s %s", vulnerabilitiesMessage, element.Description)
				}
			}
		}
	}

	if tfsecMaxSeverity == "MEDIUM" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
					vulnerabilitiesMessage = fmt.Sprintf("%s %s", vulnerabilitiesMessage, element.Description)
				}
			}
		}
	}

	if tfsecMaxSeverity == "CRITICAL" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "CRITICAL" {
					passesPolicy = false
					vulnerabilitiesMessage = fmt.Sprintf("%s %s ", vulnerabilitiesMessage, element.Description)
				}
			}
		}
	}

	if vulnerabilitiesMessage == "" {
		vulnerabilitiesMessage = "No vulnerabilities found :)"
	}

	message := fmt.Sprintf("{\"tfsec_pass\": %t, \"tfsec_max_severity\": \"%s\", \"tfsec_vulnerabilities\": \"%s\"}", passesPolicy, tfsecMaxSeverity, vulnerabilitiesMessage)

	return message
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
