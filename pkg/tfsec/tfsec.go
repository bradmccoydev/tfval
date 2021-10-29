package tfsec

import (
	"encoding/json"
	"fmt"

	"github.com/bradmccoydev/terraform-plan-validator/model"
	config "github.com/bradmccoydev/terraform-plan-validator/util"
)

func ProduceVulnerabilityReport(plan []byte) model.Vulnerabilities {
	var vulnerabilities model.Vulnerabilities
	json.Unmarshal(plan, &vulnerabilities)

	return vulnerabilities
}

func CheckIfPlanPassesTfPolicy(plan []byte, cfg config.Config) bool {
	vulnerabilities := ProduceVulnerabilityReport(plan)
	passesPolicy := true

	if IsInvalidCategory(cfg.TfsecMaxSeverity) == true {
		fmt.Println("Invalid TFSEC Severity Category defaulting to LOW")
		cfg.TfsecMaxSeverity = "LOW"
	}

	if cfg.TfsecMaxSeverity == "LOW" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "LOW" || element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
				}
			}
		}
	}

	if cfg.TfsecMaxSeverity == "MEDIUM" {
		if len(vulnerabilities.Results) > 0 {
			for _, element := range vulnerabilities.Results {
				if element.Severity == "MEDIUM" || element.Severity == "CRITICAL" {
					passesPolicy = false
				}
			}
		}
	}

	if cfg.TfsecMaxSeverity == "CRITICAL" {
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
