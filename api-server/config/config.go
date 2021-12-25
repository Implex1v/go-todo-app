package config

import (
	"go.uber.org/fx"
	"os"
)

type Config interface {
	Port() string
}

type EnvConfig struct {
	port string
}

func (c EnvConfig) Port() string {
	return c.port
}

func GetConfig() Config {
	return EnvConfig{
		port: getEnvDefault("PORT", "8080"),
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
