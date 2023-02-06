package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"strings"
)

//go:embed config.yml
var defaultConfiguration []byte

type Config struct {
	Host      string
	KafkaHost string
	Topics    []string
}

func Read(configuration *Config) error {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Configuration file
	viper.SetConfigType("yml")

	// Read configuration
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return err
	}

	// Unmarshal the configuration
	if err := viper.Unmarshal(&configuration); err != nil {
		return err
	}
	return nil
}
