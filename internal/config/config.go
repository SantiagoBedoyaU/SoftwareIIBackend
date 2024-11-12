package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type DatabaseConfig struct {
	Host     string
	DBName   string
	User     string
	Password string
}

type AuthConfig struct {
	JwtSecret string
}

type NotificationConfig struct {
	MailgunDomain string
	MailgunAPIKey string
	FrontendURL   string
}

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	Auth         AuthConfig
	Notification NotificationConfig
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env file: %v\n", err)
	}

	cfg := &Config{}
	var err error

	cfg.Server.Host = os.Getenv("SERVER_HOST")
	cfg.Server.Port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		// default port in case of error
		cfg.Server.Port = 8080
	}
	cfg.Server.ReadTimeout, _ = strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	cfg.Server.WriteTimeout, _ = strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))

	cfg.Database.Host = os.Getenv("DATABASE_HOST")
	cfg.Database.DBName = os.Getenv("DATABASE_DBNAME")
	cfg.Database.User = os.Getenv("DATABASE_USER")
	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")

	cfg.Auth.JwtSecret = os.Getenv("JWT_SECRET")

	cfg.Notification.MailgunDomain = os.Getenv("MAILGUN_DOMAIN")
	cfg.Notification.MailgunAPIKey = os.Getenv("MAILGUN_API_KEY")
	cfg.Notification.FrontendURL = os.Getenv("FRONTEND_URL")

	return cfg
}
