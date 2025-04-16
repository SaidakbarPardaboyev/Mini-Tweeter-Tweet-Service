package client

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
)

type ServiceManager interface {
}

type grpcClients struct {
}

// connect to external clients here
func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	
	return &grpcClients{}, nil
}
