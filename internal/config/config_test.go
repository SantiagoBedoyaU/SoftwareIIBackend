package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("SERVER_HOST", "testing")
	cfg := New()

	assert.Equal(t, cfg.Server.Host, "testing")
}

func TestNewConfigInvalidServerPort(t *testing.T) {
	os.Setenv("SERVER_PORT", "abc")
	cfg := New()

	assert.Equal(t, cfg.Server.Port, 8080)
}

func TestAddr(t *testing.T) {
	cfg := ServerConfig{
		Host: "localhost",
		Port: 3000,
	}

	assert.Equal(t, cfg.Addr(), "localhost:3000")
}
