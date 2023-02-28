package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Port string `env-required:"true" yaml:"port"    env:"HTTP_PORT"`
	Log  `yaml:"logger"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
}

//错误的处理凡事按这个
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}

	err = cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return cfg, nil
}
