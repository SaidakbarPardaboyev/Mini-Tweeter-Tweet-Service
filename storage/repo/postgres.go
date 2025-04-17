package repo

import (
	"context"

	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
)

type TweetsRepo interface {
	CreateTweet(ctx context.Context, in *pb.Tweet) (string, error)
	GetTweet(ctx context.Context, ID string) (*pb.Tweet, error)
	GetListTweet(ctx context.Context, in *pb.GetListTweetRequest) (*pb.TweetList, error)
	UpdateTweet(ctx context.Context, in *pb.Tweet) error
	DeleteTweet(ctx context.Context, ID string) error
}

type TweetMediaRepo interface {
	CreateTweetMedia(ctx context.Context, in *pb.TweetMedia) (string, error)
	GetTweetMedia(ctx context.Context, ID string) (*pb.TweetMedia, error)
	GetListTweetMedia(ctx context.Context, tweetID string) (*[]*pb.TweetMedia, error)
	UpdateTweetMedia(ctx context.Context, in *pb.TweetMedia) error
	DeleteTweetMedia(ctx context.Context, ID string) error
	DeleteTweetMediaWithTweetID(ctx context.Context, tweetID string) error
}
