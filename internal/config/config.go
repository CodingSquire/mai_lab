package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
type Config struct {
	//	HTTP struct {
	BindIP string `mapstructure:"HTTP_IP"`
	Port   string `mapstructure:"HTTP_PORT"`
	//	}
	//	PostgreSQL struct {
	Username string `mapstructure:"PSQL_USERNAME"`
	Password string `mapstructure:"PSQL_PASSWORD"`
	Host     string `mapstructure:"PSQL_HOST"`
	Portdb   string `mapstructure:"PSQL_PORT"`
	Database string `mapstructure:"PSQL_DATABASE"`
	//	}
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		log.Println("read application config")
		instance = &Config{}

		viper.AddConfigPath(path)
		viper.AddConfigPath(".")

		viper.SetConfigName("app") // name of config file (without extension)
		viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name

		viper.AutomaticEnv()

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {
			log.Fatalln(err)
		}
		err = viper.Unmarshal(&instance)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return instance
}
