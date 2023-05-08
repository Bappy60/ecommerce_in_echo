package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBPass     string `mapstructure:"DBPASS"`
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBPort     int    `mapstructure:"DBPORT"`
	PORT       string `mapstructure:"PORT"`
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
	REDIS_HOST string `mapstructure:"REDIS_HOST"`
	REDIS_PORT string `mapstructure:"REDIS_PORT"`
	REDIS_PASS string `mapstructure:"REDIS_PASS"`
}

func InitConfig() *Config {

	env := "prod" //os.Getenv("APP_ENV")
	viper.AddConfigPath(".")
	if env == "prod" {
		viper.SetConfigFile("prod.env")
	} else {
		viper.SetConfigFile("dev.env")
	}
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
