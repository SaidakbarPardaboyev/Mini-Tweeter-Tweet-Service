package main

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc"
	httpserver "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/http"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.Environment, "tweet_service_grpc")

	grpcServices, err := grpc.New(cfg, log)
	if err != nil {
		log.Error("Error while initializing grpcServices", logger.Error(err))
		return
	}

	httpServer, err := httpserver.New(cfg, log, grpcServices)
	if err != nil {
		log.Error("Error while initializing http server", logger.Error(err))
		return
	}

	go func() {
		err := httpServer.Run(log, cfg)
		if err != nil {
			log.Fatal("Error while running http server", logger.Error(err))
		}
	}()

	grpcServices.Run(log, cfg)
}
