package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Configure() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/statuscoin.d/")
	viper.AddConfigPath("$HOME/.statuscoin")
	viper.AddConfigPath("$HOME/.config/statuscoin")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	loadDefaults()
}

func loadDefaults() {
	viper.SetDefault("name", "SuperCoin")
	viper.SetDefault("version", "v0.0.0")
	viper.SetDefault("log.level", logrus.DebugLevel)
	viper.SetDefault("log.out", []string{"stdout", "stderr", "./logs/$date/$time.log"})
}
