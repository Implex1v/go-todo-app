package config

import (
	"go.uber.org/fx"
	"os"
)

type Config interface {
	Port() string

	DbUser() string
	DbPassword() string
	DbPort() string
	DbHost() string
	DbName() string
}

type EnvConfig struct {
	port       string
	dbUser     string
	dbPassword string
	dbPort     string
	dbHost     string
	dbName     string
}

func (c EnvConfig) DbUser() string {
	return c.dbUser
}

func (c EnvConfig) DbPassword() string {
	return c.dbPassword
}

func (c EnvConfig) DbPort() string {
	return c.dbPort
}

func (c EnvConfig) DbHost() string {
	return c.dbHost
}

func (c EnvConfig) DbName() string {
	return c.dbName
}

func (c EnvConfig) Port() string {
	return c.port
}

func GetConfig() Config {
	return &EnvConfig{
		port:       getEnvDefault("PORT", "8080"),
		dbUser:     getEnvDefault("DB_USER", "todoapp"),
		dbPassword: getEnvDefault("DB_PASSWORD", "todoapp"),
		dbPort:     getEnvDefault("DB_PORT", "5432"),
		dbHost:     getEnvDefault("DB_HOST", "localhost"),
		dbName:     getEnvDefault("DB_NAME", "todoapp"),
	}
}

func getEnvDefault(key string, defaultValue string) string {
	var value, exists = os.LookupEnv(key)

	if exists {
		return value
	} else {
		return defaultValue
	}
}

var Module = fx.Options(fx.Provide(GetConfig))
