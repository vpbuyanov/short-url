package configs

import (
	"flag"
	"os"
	"strconv"
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
		LogLevel int8
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

	logger := Logger{
		LogLevel: 1,
	}

	level, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		logLevel, err := strconv.Atoi(level)
		if err != nil {
			return nil
		}

		logger.LogLevel = int8(logLevel)
	}

	return &Config{
		Server: server,
		Logger: logger,
	}
}
