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
	server := Server{
		Address: "localhost:8080",
		BaseURL: "http://localhost:8080",
	}

	flag.StringVar(&server.Address, "a", server.Address, "input server's address")
	flag.StringVar(&server.BaseURL, "b", server.BaseURL, "input server's baseURL for shortener url")

	flag.Parse()

	addr, ok := os.LookupEnv("SERVER_ADDRESS")
	if ok {
		server.Address = addr
	}

	baseURL, ok := os.LookupEnv("BASE_URL")
	if ok {
		server.BaseURL = baseURL
	}

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
