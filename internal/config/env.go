package config

import (
	"log"

	"github.com/spf13/viper"
)

type envConfig struct {
	DbDir string `mapstructure:"LIFELINES_DB_DIR"`
	Port  int    `mapstructure:"LIFELINES_PORT"`
}

var Env *envConfig

func InitEnv() {
	Env = loadEnvVariables()
}

func loadEnvVariables() (config *envConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file: ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
