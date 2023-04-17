package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBPass     string `mapstructure:"DBPASS"`
	DBIP       string `mapstructure:"DBIP"`
	DbName     string `mapstructure:"DBNAME"`
	Port       string `mapstructure:"PORT"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

func InitConfig() *Config {

	viper.AddConfigPath("D:/goProjects/ecommerce_in_echo")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return config

}

func SetConfig() {
	LocalConfig = InitConfig()
}
