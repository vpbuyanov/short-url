package services

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/configs"
)

func InitLogger(cfg configs.Logger) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	err := logger.Level.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal log level: %w", err)
	}

	return logger, nil
}
