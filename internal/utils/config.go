package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	Spotify struct {
		ClientID      string `mapstructure:"clientId"`
		ClientSecret  string `mapstructure:"clientSecret"`
		TokenLocation string `mapstructure:"tokenLocation"`
	}
	Server struct {
		Host      string `mapstructure:"host"`
		Port      int    `mapstructure:"port"`
		QueueSize int    `mapstructure:"queueSize"`
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

func (c *Config) ReadToken() (*oauth2.Token, error) {
	if c.Spotify.TokenLocation == "database" {
		return nil, fmt.Errorf("token location is set to database")
	}

	if c.Spotify.TokenLocation == "" {
		return nil, fmt.Errorf("token location is not set")
	}

	return readTokenFromFile(c.Spotify.TokenLocation)
}

func readTokenFromFile(tokenLocation string) (*oauth2.Token, error) {
	data, err := os.ReadFile(tokenLocation)
	if err != nil {
		return nil, fmt.Errorf("failed reading token file: %v", err)
	}

	var token oauth2.Token

	err = json.Unmarshal(data, &token)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshalling token: %v", err)
	}

	return &token, nil
}
