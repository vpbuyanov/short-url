package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/handlers"
)

type server struct {
	cfg *configs.Server
}

type Server interface {
	Start(context.Context, *logrus.Logger)
}

func New(config *configs.Server) Server {
	return &server{
		cfg: config,
	}
}

func (s *server) Start(ctx context.Context, log *logrus.Logger) {
	serv := fiber.New()
	serv.Use(logger.New())

	h := handlers.New(log, s.cfg)
	h.RegisterRouter(serv)

	err := serv.Listen(s.cfg.Address)
	if err != nil {
		return
	}
}
