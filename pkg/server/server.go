package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/doffy007/golang-hashmap/config"
	"github.com/doffy007/golang-hashmap/internal/logger"
	"github.com/doffy007/golang-hashmap/internal/routers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	config config.Config
}

func NewServer() (*Server, error) {
	cnfg, err := config.NewParsedConfig()
	if err != nil {
		return nil, err
	}

	log := logger.NewLogger()
	router := mux.NewRouter()
	routers.Register(router, log)

	s := Server{
		logger: log,
		router: router,
		config: cnfg,
	}

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.ServerPort),
		Handler: s.router,
	}

	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	defer signal.Stop(stopServer)

	var eg errgroup.Group

	eg.Go(func() error {
		s.logger.Printf("API listening on port:%d", s.config.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("error: starting API server: %w", err)
		}
		return nil
	})

	shutdownComplete := make(chan struct{})

	eg.Go(func() error {
		defer close(shutdownComplete)
		<-stopServer
		s.logger.Warn("server receive stop signal")
		err := server.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("graceful shutdown no complete: %w", err)
		}
		s.logger.Info("server was shut down gracefully")
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil

}
