package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	OpaGcpPolicy   string
	OpaAzurePolicy string
	OpaAwsPolicy   string
}

func LoadConfig() (*Config, error) {
	var conf = &Config{}

	v := viper.New()

	viper.AutomaticEnv()

	v.SetConfigName("conf")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	v.SetEnvPrefix("VTT")
	v.AutomaticEnv()

	v.SetDefault("OpaGcpPolicy", "opa-gcp-policy.rego")
	v.SetDefault("OpaAzurePolicy", "opa-azure-policy.rego")
	v.SetDefault("OpaAwsPolicy", "opa-aws-policy.rego")

	err := v.ReadInConfig()

	if err != nil {
		return nil, err
	}

	conf.OpaGcpPolicy = strings.TrimSpace(v.GetString("OpaGcpPolicy"))
	conf.OpaAzurePolicy = strings.TrimSpace(v.GetString("OpaAzurePolicy"))
	conf.OpaAwsPolicy = strings.TrimSpace(v.GetString("OpaAwsPolicy"))

	return conf, nil
}
