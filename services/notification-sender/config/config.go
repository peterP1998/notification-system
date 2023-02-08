package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"log"
	"strings"
	"fmt"
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

	log.Print(configuration)
	return nil
}

func GetConfigProperty(property string) string {
	return fmt.Sprintf("%v",viper.Get(property))
}
