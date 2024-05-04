package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/handlers"
)

type server struct {
	log *logrus.Logger
	cfg configs.Server
}

type Server interface {
	Start(ctx context.Context)
}

func New(log *logrus.Logger, config configs.Server) Server {
	return &server{
		log: log,
		cfg: config,
	}
}

func (s *server) Start(ctx context.Context) {
	serv := fiber.New()
	serv.Use(logger.New())

	address := fmt.Sprintf("%v:%v", s.cfg.Host, s.cfg.Port)

	h := handlers.New(s.log)
	h.RegisterRouter(serv)

	err := serv.Listen(address)
	if err != nil {
		s.log.Errorf("Error starting server, err: %v", err)
		return
	}
}
