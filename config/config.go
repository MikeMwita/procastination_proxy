package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {

	//database configuration
	DbUser    string        `default:"admin"`
	DbPass    string        `default:"admin"`
	DbHost    string        `default:"0.0.0.0"`
	DbPort    string        `default:"3306"`
	DbName    string        `default:"admin"`
	DbTimeout time.Duration `default:"10s"`
	//server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
}

var Cfg config

func LoadConfig() (*Config, error) {
	jwtExpiry, ok := os.LookupEnv("JWT_EXPIRY")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable JWT_EXPIRY")
	}
	cfg := &Config{
		Jwt: Jwt{
			ExpiryMinutes:     jwtExpiryInt,
			RefreshExpiryDays: jwtRefreshExpiryInt,
		},
		Database: DatabaseService{
			Port: dbPort,
			Host: dbHost,
		},
	}
	return Cfg, nil
}
