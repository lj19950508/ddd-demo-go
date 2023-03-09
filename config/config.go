package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)


//分清哪些是必须 哪些可以默认
type Config struct {
	HttpServer HttpServer `yaml:"httpserver"`
	Log Log   `yaml:"logger"`
	Mysql Mysql `yaml:"mysql"`
}

type HttpServer struct{
	Port  string `env-required:"true" yaml:"port"    env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
}

type Mysql struct {
	Url string `env-required:"true" yaml:"url" env:"MYSQL_URL"`
}

func New() (cfg *Config) {
	cfg = &Config{}
	if err := cleanenv.ReadConfig("config/config.yml", cfg); err != nil {
		log.Fatalf("%s",err)
	}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatalf("%s",err)
	}
	return cfg
}
