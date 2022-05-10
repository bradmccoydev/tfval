package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the app
// The values are read by viper from a config file or env vars
type Config struct {
	OpaRegoQuery     string `mapstructure:"OPA_REGO_QUERY"`
	TfsecMaxSeverity string `mapstructure:"TFSEC_MAX_SEVERITY"`
	RepoBaseURL      string `mapstructure:"REPO_BASEURL"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return
}
