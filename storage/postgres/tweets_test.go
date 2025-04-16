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

func Init() (repo.TweetsRepo, error) {
	cfg := config.Load()

	fmt.Println(cfg.PostgresPassword)

	log := logger.New(cfg.Environment, "tweets_service_grpc")
	psql, err := db.NewPostgresDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	tweetRepo := NewTweetsRepo(psql, log, cfg)

	return tweetRepo, nil
}

func TestCreateTweet(t *testing.T) {
	tweetsRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := tweetsRepo.CreateTweet(context.Background(), &pb.Tweet{
		Content:   "Saidakbar for test",
		UserId:    "76c50a9f-8e3f-40d5-bf5b-efbd33084947",
		Medias:    []*pb.Media{},
		CreatedAt: "",
		UpdatedAt: "",
	})
	if err != nil {
		t.Error(fmt.Errorf("error creating tweet: %v", err))
	}

	fmt.Println("ID: ", resp)
}

func TestGetTweet(t *testing.T) {
	tweetRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	resp, err := tweetRepo.GetTweet(context.Background(), "2020f6ac-73bf-4593-90ed-da3f71ccf814")
	if err != nil {
		t.Error(fmt.Errorf("error creating tweet: %v", err))
	}

	fmt.Println("tweet info: ", resp)
}

func TestGetListTweet(t *testing.T) {
	tweetRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	var search = ""
	resp, err := tweetRepo.GetListTweet(context.Background(), &pb.GetListTweetRequest{
		Search: &search,
		Page:   1,
		Limit:  10,
	})
	if err != nil {
		t.Error(fmt.Errorf("getting list of tweets: %v", err))
	}

	for _, tweet := range resp.Tweets {
		fmt.Println("Tweet Id: ", tweet.Id)
		fmt.Println("Tweet Content: ", tweet.Content)
		fmt.Println("Tweet UserId: ", tweet.UserId)
		fmt.Println("Tweet CreatedAt: ", tweet.CreatedAt)
		fmt.Println("Tweet UpdatedAt: ", tweet.UpdatedAt)
		for _, media := range tweet.Medias {
			fmt.Println("Tweet Media Id: ", media.Id)
			fmt.Println("Tweet Media TweetId: ", media.TweetId)
			fmt.Println("Tweet Media Url: ", media.Url)
		}
	}
}

func TestUpdateTweet(t *testing.T) {
	tweetRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = tweetRepo.UpdateTweet(context.Background(), &pb.Tweet{
		Id:      "60800c51-03b3-4184-87bb-3daf82a2145d",
		Content: "Saidakbar for test for for",
		UserId:  "76c50a9f-8e3f-40d5-bf5b-efbd33084947",
		Medias:  []*pb.Media{},
	})

	if err != nil {
		t.Error(fmt.Errorf("error updating tweet: %v", err))
	}
}

func TestDeleteTweet(t *testing.T) {
	tweetRepo, err := Init()
	if err != nil {
		t.Error(fmt.Errorf("error while connecting to database: %v", err))
	}

	err = tweetRepo.DeleteTweet(context.Background(), "2020f6ac-73bf-4593-90ed-da3f71ccf814")
	if err != nil {
		t.Error(fmt.Errorf("error deleting tweet: %v", err))
	}
}
