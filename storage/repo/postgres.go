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
