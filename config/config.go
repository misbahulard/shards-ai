package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/opt/shards-ai")
	viper.AddConfigPath("$HOME/.shards-ai")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
