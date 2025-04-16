package services

import (
	"context"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"
)

type TweetsService struct {
	storage  storage.StorageI
	logger   logger.Logger
	cfg      *config.Config
	services client.ServiceManager
	pb.UnimplementedTweetsServiceServer
}

func NewUsersService(options *ServiceOptions) pb.TweetsServiceServer {
	return &TweetsService{
		storage:  options.Storage,
		logger:   options.Logger,
		services: options.ServiceManager,
		cfg:      options.Config,
	}
}

func (s *TweetsService) CreateTweet(ctx context.Context, in *pb.CreateTweetRequest) (*pb.CreateTweetResponse, error) {
	tweetID, err := s.storage.Tweets().CreateTweet(ctx, in.Tweet)
	if err != nil {
		return nil, err
	}

	for _, media := range in.Tweet.Medias {
		media.TweetId = tweetID
		_, err := s.storage.TweetMedias().CreateTweetMedia(ctx, media)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateTweetResponse{Id: tweetID}, nil
}

func (s *TweetsService) GetTweet(ctx context.Context, in *pb.GetTweetRequest) (*pb.Tweet, error) {
	tweetResp, err := s.storage.Tweets().GetTweet(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	tweetMedias, err := s.storage.TweetMedias().GetListTweetMedia(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	tweetResp.Medias = *tweetMedias

	return tweetResp, nil
}

func (s *TweetsService) GetListTweet(ctx context.Context, in *pb.GetListTweetRequest) (*pb.TweetList, error) {
	if in.Limit == 0 {
		in.Limit = 10
	}

	if in.Page == 0 {
		in.Page = 1
	}

	return s.storage.Tweets().GetListTweet(ctx, in)
}

func (s *TweetsService) UpdateTweet(ctx context.Context, in *pb.UpdateTweetRequest) (*pb.UpdateTweetResponse, error) {
	err := s.storage.Tweets().UpdateTweet(ctx, in.Tweet)
	if err != nil {
		return &pb.UpdateTweetResponse{Success: false}, err
	}

	return &pb.UpdateTweetResponse{Success: true}, nil
}

func (s *TweetsService) DeleteTweet(ctx context.Context, in *pb.DeleteTweetRequest) (*pb.DeleteTweetResponse, error) {
	err := s.storage.Tweets().DeleteTweet(ctx, in.Id)
	if err != nil {
		return &pb.DeleteTweetResponse{Success: false}, err
	}

	err = s.storage.TweetMedias().DeleteTweetMediaWithTweetID(ctx, in.Id)
	if err != nil {
		return &pb.DeleteTweetResponse{Success: false}, err
	}

	return &pb.DeleteTweetResponse{Success: true}, nil
}
