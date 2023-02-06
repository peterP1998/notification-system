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
}

func Read(configuration *Config) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return err
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return err
	}
	return nil
}
