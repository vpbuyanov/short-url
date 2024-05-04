package configs

import "os"

type (
	Config struct {
		Server Server
		Logger Logger
	}

	Server struct {
		Host string
		Port string
	}

	Logger struct {
		LogLevel string
	}
)

func LoadConfig() *Config {
	serverPort := os.Getenv("SERVER_PORT")
	serverHost := os.Getenv("SERVER_HOST")

	if serverPort == "" {
		serverPort = "8080"
	}

	if serverHost == "" {
		serverHost = "localhost"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		Server: Server{
			Host: serverHost,
			Port: serverPort,
		},
		Logger: Logger{
			LogLevel: logLevel,
		},
	}
}
