package internal

import (
	"github.com/archey347/go-http-api-template/internal/http"
	"github.com/spf13/viper"
)

type Config struct {
	Http http.Config `mapstructure:"http"`
}

const defaultConfigFile = "/etc/go-http-api-template/go-http-api-template.yaml"

func LoadConfig(configFile string) (*Config, error) {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
