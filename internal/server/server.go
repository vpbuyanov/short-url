package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/handlers"
	"github.com/vpbuyanov/short-url/internal/repos"
)

type Server struct {
	cfg *configs.Server
}

func New(config *configs.Server) Server {
	return Server{
		cfg: config,
	}
}

func (s *Server) Start(ctx context.Context, log *logrus.Logger, url repos.URL) error {
	serv := fiber.New()
	serv.Use(logger.New())

	h := handlers.New(log, s.cfg, url)
	h.RegisterRouter(serv)

	err := serv.Listen(s.cfg.Address)
	if err != nil {
		return fmt.Errorf("can't lestener port, err: %w", err)
	}

	return nil
}
