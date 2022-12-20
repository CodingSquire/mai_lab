package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	isDebug bool `yaml:"is_debug" env-required:"true"`
	HTTP    struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"http"`
	PostgreSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalln(err)
		}
	})
	return instance
}
