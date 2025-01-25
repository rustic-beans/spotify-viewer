package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Spotify struct {
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
		TokenFile    string `yaml:"tokenFile"`
	}
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver string `yaml:"driver"`
		Source string `yaml:"source"`
	}
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := NewConfig()

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetURL() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
