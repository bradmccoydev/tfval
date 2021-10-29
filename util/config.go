package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the app
// The values are read by viper from a config file or env vars
type Config struct {
	OpaGcpPolicy     string `mapstructure:"OPA_GCP_POLICY"`
	OpaAzurePolicy   string `mapstructure:"OPA_AZURE_POLICY"`
	OpaAwsPolicy     string `mapstructure:"OPA_AWS_POLICY"`
	OpaRegoQuery     string `mapstructure:"OPA_REGO_QUERY"`
	TfsecMaxSeverity string `mapstructure:"TFSEC_MAX_SEVERITY"`
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
