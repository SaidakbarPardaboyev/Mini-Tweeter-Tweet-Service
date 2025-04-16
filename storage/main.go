package storage

import (
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/postgres"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
)

type StorageI interface {
	Tweets() repo.TweetsRepo
}

type storagePg struct {
	tweets repo.TweetsRepo
}

func New(db *db.Postgres, log logger.Logger, cfg *config.Config) StorageI {
	return &storagePg{
		tweets: postgres.NewTweetsRepo(db, log, cfg),
	}
}

func (s *storagePg) Tweets() repo.TweetsRepo {
	return s.tweets
}