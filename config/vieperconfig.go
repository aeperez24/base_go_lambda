package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadViperConfig(path string, environment string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(environment)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	keys := viper.AllKeys()
	for _, key := range keys {
		uppercaseKey := strings.ToUpper(key)
		os.Setenv(uppercaseKey, viper.GetString(strings.ToUpper(key)))

	}
}
