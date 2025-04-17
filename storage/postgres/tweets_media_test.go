package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
)

func InitTweetMedia() (repo.TweetMediaRepo, error) {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "tweets_service_grpc")
	psql, err := db.NewPostgresDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	tweetMediaRepo := NewTweetMediaRepo(psql, log, cfg)

	return tweetMediaRepo, nil
}

func TestCreateTweetMedia(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := tweetMediaRepo.CreateTweetMedia(context.Background(), &pb.TweetMedia{
		TweetId:   "60800c51-03b3-4184-87bb-3daf82a2145d",
		MediaType: "video",
		Url:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating tweet media: %v", err))
	}

	fmt.Println("ID: ", resp)
}

func TestGetTweetMedia(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := tweetMediaRepo.GetTweetMedia(context.Background(), "c32e6e62-af51-4bc9-b3b1-64dc9ce6bcdf")
	if err != nil {
		t.Error(fmt.Errorf("error getting tweet media: %v", err))
	}

	fmt.Println("tweet media info: ", resp)
}

func TestGetListTweetMedia(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := tweetMediaRepo.GetListTweetMedia(context.Background(), "60800c51-03b3-4184-87bb-3daf82a2145d")
	if err != nil {
		t.Error(fmt.Errorf("getting list of tweet media: %v", err))
	}

	for _, media := range *resp {
		fmt.Println("Tweet Media Id: ", media.Id)
		fmt.Println("Tweet Media TweetId: ", media.TweetId)
		fmt.Println("Tweet Media Url: ", media.Url)
		fmt.Println("Tweet Media MediaType: ", media.MediaType)
	}
}

func TestUpdateTweetMedia(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = tweetMediaRepo.UpdateTweetMedia(context.Background(), &pb.TweetMedia{
		Id:        "c32e6e62-af51-4bc9-b3b1-64dc9ce6bcdf",
		TweetId:   "60800c51-03b3-4184-87bb-3daf82a2145d",
		MediaType: "image",
		Url:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ_test",
	})

	if err != nil {
		t.Error(fmt.Errorf("error updating tweet media: %v", err))
	}
}

func TestDeleteTweetMedia(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = tweetMediaRepo.DeleteTweetMedia(context.Background(), "c32e6e62-af51-4bc9-b3b1-64dc9ce6bcdf")
	if err != nil {
		t.Error(fmt.Errorf("error deleting tweet media: %v", err))
	}
}

func TestDeleteTweetMediaWithTweetID(t *testing.T) {
	tweetMediaRepo, err := InitTweetMedia()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = tweetMediaRepo.DeleteTweetMediaWithTweetID(context.Background(), "60800c51-03b3-4184-87bb-3daf82a2145d")
	if err != nil {
		t.Error(fmt.Errorf("error deleting tweet media with tweet id: %v", err))
	}
}
