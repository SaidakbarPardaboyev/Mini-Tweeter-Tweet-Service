package services

import (
	"fmt"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/users_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"
)

type ServiceOptions struct {
	ServiceManager client.ServiceManager
	Storage        storage.StorageI
	Config         *config.Config
	Logger         logger.Logger
}

func validatingUser(in *pb.User) (bool, error) {
	fmt.Println(in.Name)

	if in.Name == "" || len(in.Name) < 3 {
		return false, fmt.Errorf("name must be at least 3 characters")
	}

	if in.Username == "" || len(in.Username) < 8 {
		return false, fmt.Errorf("username must be at least 8 characters")
	}

	if in.PasswordHash == "" {
		return false, fmt.Errorf("invalid password")
	}

	return true, nil
}
