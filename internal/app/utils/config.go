package utils

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var (
	configOnce sync.Once
	config     Config
)

type Config struct {
	Debug         bool
	DebugSQL      bool
	CORS          bool
	GZIP          bool
	SessionKey    string
	ServerAddress string
	DBDriver      string
	DBString      string
	StaticPath    string
	LogPath       string
}

func GetConfig() Config {
	configOnce.Do(func() {
		if os.Getenv("DEBUG") == "1" {
			config.Debug = true
		}

		if os.Getenv("DEBUG_SQL") == "1" {
			config.DebugSQL = true
		}

		if os.Getenv("CORS") == "1" {
			config.CORS = true
		}

		if os.Getenv("GZIP") == "1" {
			config.GZIP = true
		}

		if config.SessionKey = os.Getenv("SESSION_KEY"); config.SessionKey == "" {
			panic("config: SESSION_KEY variable is required but not set")
		}

		if config.ServerAddress = os.Getenv("SERVER_ADDRESS"); config.ServerAddress == "" {
			panic("config: SERVER_ADDRESS variable is required but not set")
		}

		if config.DBDriver = os.Getenv("DB_DRIVER"); config.DBDriver == "" {
			panic("config: DB_DRIVER variable is required but not set")
		}

		if config.DBString = os.Getenv("DB_STRING"); config.DBString == "" {
			panic("config: DB_STRING variable is required but not set")
		}

		if config.StaticPath = os.Getenv("STATIC_PATH"); config.StaticPath != "" {
			os.MkdirAll(config.StaticPath, os.ModePerm)
		}

		if config.LogPath = os.Getenv("LOG_PATH"); config.LogPath != "" {
			os.MkdirAll(config.LogPath, os.ModePerm)
		}
	})

	return config
}
