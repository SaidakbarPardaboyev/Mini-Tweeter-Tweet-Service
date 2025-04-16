package handler

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc"
)

type Optoins struct {
	Service *grpc.GRPCService
	Logger  logger.Logger
	Cfg     *config.Config
}

type Handler struct {
	service *grpc.GRPCService
	logger  logger.Logger
	cfg     *config.Config
}

func NewHandler(options Optoins) *Handler {
	return &Handler{
		service: options.Service,
		logger:  options.Logger,
		cfg:     options.Cfg,
	}
}
