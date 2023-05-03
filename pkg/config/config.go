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
	PORT       string `mapstructure:"PORT"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
	REDIS_HOST string `mapstructure:"REDIS_HOST"`
	REDIS_PORT string `mapstructure:"REDIS_PORT"`
	REDIS_PASS string `mapstructure:"REDIS_PASS"`
}

func InitConfig() *Config {

	viper.AddConfigPath("D:/goProjects/ecommerce_in_echo/")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file x ", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file while unmarshaling", err)
	}

	return config

}

func SetConfig() {
	LocalConfig = InitConfig()
}
