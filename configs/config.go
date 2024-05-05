package configs

import (
	"flag"
	"os"
)

type (
	Config struct {
		Server Server
		Logger Logger
	}

	Server struct {
		Address string
		BaseURL string
	}

	Logger struct {
		LogLevel string
	}
)

func LoadConfig() *Config {
	server := Server{}

	server.Address = os.Getenv("SERVER_ADDRESS")
	server.BaseURL = os.Getenv("BASE_URL")

	if server.Address == "" {
		flag.StringVar(&server.Address, "a", "localhost:8080", "input server's address")
	}

	if server.BaseURL == "" {
		flag.StringVar(&server.BaseURL, "b", "http://localhost:8080", "input server's port")
	}

	flag.Parse()

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		Server: server,
		Logger: Logger{
			LogLevel: logLevel,
		},
	}
}
