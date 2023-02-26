package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int
}

//错误的处理凡事按这个
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}

	err = cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		return
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return
	}
	return cfg, nil
}
