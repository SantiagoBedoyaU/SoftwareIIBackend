package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}

	cfg := &Config{}

	cfg.Server.Host = os.Getenv("server.host")
	cfg.Server.Port, _ = strconv.Atoi(os.Getenv("server.host"))
	cfg.Server.ReadTimeout, _ = strconv.Atoi(os.Getenv("server.read_timeout"))
	cfg.Server.WriteTimeout, _ = strconv.Atoi(os.Getenv("server.write_timeout"))

	cfg.Database.Host = os.Getenv("database.host")
	cfg.Database.DBName = os.Getenv("database.dbname")
	cfg.Database.User = os.Getenv("database.user")
	cfg.Database.Password = os.Getenv("database.password")

	return cfg
}
