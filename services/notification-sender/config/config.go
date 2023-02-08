package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

//go:embed config.yml
var defaultConfiguration []byte

type Config struct {
	Host        string
	KafkaHost   string
	Topics      []string
	EmailConfig Email
	SMSConfig   SMS
}

type Email struct {
	From     string
	Addr     string
	Host     string
	Password string
}

type SMS struct {
	AccountId    string
	AccountToken string
	From         string
}

func Read(configuration *Config) error {
	log.Print("Loading Configuration")

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

	log.Print("Loading configuration finished")

	return nil
}

func GetConfigProperty(property string) string {
	return fmt.Sprintf("%v", viper.Get(property))
}
