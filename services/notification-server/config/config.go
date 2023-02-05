package config

import (
	"github.com/spf13/viper"
	"strings"
	"bytes"
	_ "embed"
)

//go:embed config.yml
var defaultConfiguration []byte

type Config struct {
	Host     string
	KafkaHost     string
}

var Configuration *Config

func Read() (error) {
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
	if err := viper.Unmarshal(&Configuration); err != nil {
		return err
	}
	return nil
}