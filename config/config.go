package config

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
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
func New(lc fx.Lifecycle) (cfg *Config) {
	cfg = &Config{}
	lc.Append(fx.Hook{
		//被需要的时候只会执行一次
		OnStart: func(ctx context.Context) error {
			fmt.Println("start")
			if err := cleanenv.ReadConfig("config/config.yml", cfg); err != nil {
				return err
			}
			if err := cleanenv.ReadEnv(cfg); err != nil {
				return err
			}
			return nil
			// return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("stop-------")
			return nil
		},
	})
	return cfg
}
