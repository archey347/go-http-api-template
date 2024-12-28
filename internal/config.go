package internal

import (
	"github.com/archey347/dynamic-dns/dynamic-dns/internal/http"
	"github.com/spf13/viper"
)

type Config struct {
	Http    http.Config       `mapstructure:"http"`
	Keys    map[string]string `mapstructure:"keys"`
	Zones   []Zone            `mapstructure:"zones"`
	Updates []Update          `mapstructure:"updates"`
}

type Zone struct {
	Name    string   `mapstructure:"name"`
	Pattern string   `mapstructure:"pattern"`
	Keys    []string `mapstructure:"keys"`
}

type Update struct {
	Address string `mapstructure:"address"`
	Secret  string `mapstructure:"secret"`
}

const defaultConfigFile = "/etc/dynamic-dns/dynamic-dns-server.yaml"

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
