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
	ReadOnly      bool
	AdminLogin    string
	AdminPassword string
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

		if os.Getenv("READ_ONLY") == "1" {
			config.ReadOnly = true
		}

		if config.AdminLogin = os.Getenv("ADMIN_LOGIN"); config.AdminLogin == "" {
			panic("config: ADMIN_LOGIN variable is required but not set")
		}

		if config.AdminPassword = os.Getenv("ADMIN_PASSWORD"); config.AdminPassword == "" {
			panic("config: ADMIN_PASSWORD variable is required but not set")
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
