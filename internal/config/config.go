package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Debug         bool
	ReadOnly      bool
	ServerAddress string
	DBDriver      string
	DBString      string
}

func GetConfig() Config {
	config := Config{
		Debug:         getenv("DEBUG", false) == "1",
		ReadOnly:      getenv("READ_ONLY", false) == "1",
		ServerAddress: getenv("SERVER_ADDRESS", true),
		DBDriver:      getenv("DB_DRIVER", true),
		DBString:      getenv("DB_STRING", true),
	}
	return config
}

func getenv(key string, required bool) string {
	value := os.Getenv(key)
	if value == "" && required {
		panic(fmt.Sprintf("config: %s variable is required but not set", key))
	}
	return value
}
