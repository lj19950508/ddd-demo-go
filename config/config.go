package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port  string `env-required:"true" yaml:"port"    env:"HTTP_PORT"`
	Log   `yaml:"logger"`
	Mysql `yaml:"mysql"`
}

type Log struct {
	Level string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
}

type Mysql struct {
	Url string `env-required:"true" yaml:"url" env:"MYSQL_URL"`
}

//这么使用指针会造成值拷贝
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
