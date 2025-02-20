package utils

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Spotify struct {
		ClientID     string `mapstructure:"clientId"`
		ClientSecret string `mapstructure:"clientSecret"`
		TokenFile    string `mapstructure:"tokenFile"`
	}
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	}
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

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
