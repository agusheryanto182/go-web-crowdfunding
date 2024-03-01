package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joeshaw/envdecode"
)

type Config struct {
	AppPort  int `env:"PORT,required"`
	Database database
	Secret   string `env:"SECRET,required"`
	// ClientKeyMidtrans string `env:"CLIENT_MIDTRANS,required"`
	// ServerKeyMidtrans string `env:"SERVER_MIDTRANS,required"`
}

type database struct {
	DbHost string `env:"DB_HOST,required"`
	DbPort string `env:"DB_PORT,required"`
	DbUser string `env:"DB_USER,required"`
	DbPass string `env:"DB_PASS,required"`
	DbName string `env:"DB_NAME,required"`
}

func NewConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
