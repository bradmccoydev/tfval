package model

type TfValidation []struct {
	ValidationPassed   bool `json:"validation_passed"`
	Score              int  `json:"score"`
	MaxAcceptableScore int  `json:"max_acceptable_score"`
	Data               struct {
		Address string `json:"address"`
		Change  struct {
			Actions        []string    `json:"actions"`
			After          interface{} `json:"after"`
			AfterSensitive bool        `json:"after_sensitive"`
			AfterUnknown   struct {
			} `json:"after_unknown"`
			Before struct {
				ID       string `json:"id"`
				Location string `json:"location"`
				Name     string `json:"name"`
				Tags     struct {
				} `json:"tags"`
				Timeouts interface{} `json:"timeouts"`
			} `json:"before"`
			BeforeSensitive struct {
				Tags struct {
				} `json:"tags"`
			} `json:"before_sensitive"`
		} `json:"change"`
		Mode         string `json:"mode"`
		Name         string `json:"name"`
		ProviderName string `json:"provider_name"`
		Type         string `json:"type"`
	} `json:"data"`
}
