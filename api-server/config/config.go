package config

import "os"

type Config struct {
	Port string
}

func GetConfig() *Config {
	return &Config{
		Port: getEnvDefault("PORT", "8080"),
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
