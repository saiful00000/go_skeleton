package configs

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DbUser string `mapstructure:"DBUSER"`
	DbPass string `mapstructure:"DBPASS"`
	DbIp   string `mapstructure:"DBIP"`
	DbName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
}

func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadConfig(); err != nil {
		log.Fatal("Error while reading app.env file", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading app.env file", err)
	}

	return config
}

func SetConfig() {
	LocalConfig = InitConfig()
}
