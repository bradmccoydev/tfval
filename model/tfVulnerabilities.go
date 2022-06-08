package model

type Vulnerabilities struct {
	Results []struct {
		RuleID          string   `json:"rule_id"`
		LegacyRuleID    string   `json:"legacy_rule_id"`
		RuleDescription string   `json:"rule_description"`
		RuleProvider    string   `json:"rule_provider"`
		Impact          string   `json:"impact"`
		Resolution      string   `json:"resolution"`
		Links           []string `json:"links"`
		Description     string   `json:"description"`
		Severity        string   `json:"severity"`
		Status          string   `json:"status"`
		Location        struct {
			Filename  string `json:"filename"`
			StartLine int    `json:"start_line"`
			EndLine   int    `json:"end_line"`
		} `json:"location"`
	} `json:"results"`
}
