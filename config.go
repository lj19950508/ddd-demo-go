package main

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
}

func NewConfig() (cfg Config, err error) {
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		return
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return
	}
	return cfg, nil
}
