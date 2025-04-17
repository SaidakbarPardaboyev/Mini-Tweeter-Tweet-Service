package services

import (
	"context"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/server/grpc/client"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage"
)

type TweetMediasService struct {
	storage  storage.StorageI
	logger   logger.Logger
	cfg      *config.Config
	services client.ServiceManager
	pb.UnimplementedTweetMediasServiceServer
}

func NewTweetMediasService(options *ServiceOptions) pb.TweetMediasServiceServer {
	return &TweetMediasService{
		storage:  options.Storage,
		logger:   options.Logger,
		services: options.ServiceManager,
		cfg:      options.Config,
	}
}

func (t *TweetMediasService) CreateTweetMedia(ctx context.Context, in *pb.CreateTweetMediaRequest) (*pb.CreateTweetMediaResponse, error) {
	tweetMediaID, err := t.storage.TweetMedias().CreateTweetMedia(ctx, in.TweetMedia)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTweetMediaResponse{Id: tweetMediaID}, nil
}

func (t *TweetMediasService) GetTweetMedia(ctx context.Context, in *pb.GetTweetMediaRequest) (*pb.TweetMedia, error) {
	tweetMediaResp, err := t.storage.TweetMedias().GetTweetMedia(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	return tweetMediaResp, nil
}

func (t *TweetMediasService) GetListTweetMedia(ctx context.Context, in *pb.GetListTweetMediaRequest) (*pb.TweetMediaList, error) {
	tweetMedias, err := t.storage.TweetMedias().GetListTweetMedia(ctx, in.GetTweetId())
	if err != nil {
		return nil, err
	}

	return &pb.TweetMediaList{
		TweetMedias: *tweetMedias,
		Count:       int64(len(*tweetMedias)),
	}, nil
}

func (t *TweetMediasService) UpdateTweetMedia(ctx context.Context, in *pb.UpdateTweetMediaRequest) (*pb.UpdateTweetMediaResponse, error) {
	err := t.storage.TweetMedias().UpdateTweetMedia(ctx, in.TweetMedia)
	if err != nil {
		return &pb.UpdateTweetMediaResponse{Success: false}, err
	}

	return &pb.UpdateTweetMediaResponse{Success: true}, nil
}

func (t *TweetMediasService) DeleteTweetMedia(ctx context.Context, in *pb.DeleteTweetMediaRequest) (*pb.DeleteTweetMediaResponse, error) {
	err := t.storage.TweetMedias().DeleteTweetMedia(ctx, in.Id)
	if err != nil {
		return &pb.DeleteTweetMediaResponse{Success: false}, err
	}

	return &pb.DeleteTweetMediaResponse{Success: true}, nil
}

func (t *TweetMediasService) DeleteTweetMediaWithTweetID(ctx context.Context, in *pb.DeleteTweetMediaWithTweetIDRequest) (*pb.DeleteTweetMediaWithTweetIDResponse, error) {
	err := t.storage.TweetMedias().DeleteTweetMedia(ctx, in.TweetId)
	if err != nil {
		return &pb.DeleteTweetMediaWithTweetIDResponse{Success: false}, err
	}

	return &pb.DeleteTweetMediaWithTweetIDResponse{Success: true}, nil
}
