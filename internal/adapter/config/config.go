package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

func New() (*Config, error) {
	var cfg Config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
