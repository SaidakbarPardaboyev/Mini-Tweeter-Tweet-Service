package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/static"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
	"github.com/google/uuid"
)

type tweetMedia struct {
	db  *db.Postgres
	cff *config.Config
	log logger.Logger
}

func NewTweetMediaRepo(db *db.Postgres, log logger.Logger, cff *config.Config) repo.TweetMediaRepo {
	return &tweetMedia{db: db, log: log, cff: cff}
}

func (t *tweetMedia) CreateTweetMedia(ctx context.Context, in *pb.Media) (string, error) {
	var (
		id = uuid.New().String()

		// Insert the new tweet media into the database
		insertQuery = static.InsertTweetMediaQuery
	)

	_, err := t.db.Db.Exec(ctx, insertQuery,
		id,
		in.TweetId,
		in.MediaType,
		in.Url,
	)
	if err != nil {
		return "", fmt.Errorf("error creating tweet media: %v", err)
	}

	return id, nil
}

func (t *tweetMedia) GetTweetMedia(ctx context.Context, ID string) (*pb.Media, error) {
	query := static.GetTweetMediaByIDQuery

	var (
		tweetMedia pb.Media
	)

	err := t.db.Db.QueryRow(ctx, query, ID).Scan(
		&tweetMedia.Id,
		&tweetMedia.TweetId,
		&tweetMedia.MediaType,
		&tweetMedia.Url,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tweet media not found")
		}
		return nil, fmt.Errorf("error getting tweet media: %v", err)
	}

	return &tweetMedia, nil
}

func (t *tweetMedia) GetListTweetMedia(ctx context.Context, tweetID string) (*[]*pb.Media, error) {
	var (
		args      = []interface{}{tweetID}
		queryBase = static.GetListTweetMediaBaseQuery
	)

	rows, err := t.db.Db.Query(ctx, queryBase, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching tweet media list: %v", err)
	}
	defer rows.Close()

	var tweetMedias []*pb.Media
	for rows.Next() {
		var (
			tweetMedia pb.Media
		)

		err := rows.Scan(
			&tweetMedia.Id,
			&tweetMedia.TweetId,
			&tweetMedia.MediaType,
			&tweetMedia.Url,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning tweet media: %v", err)
		}

		tweetMedias = append(tweetMedias, &tweetMedia)
	}

	// Check for row iteration errors
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tweet medias: %v", err)
	}

	return &tweetMedias, nil
}

func (t *tweetMedia) UpdateTweetMedia(ctx context.Context, in *pb.Media) error {
	// Query to update tweet details
	query := static.UpdateTweetMediaQuery

	_, err := t.db.Db.Exec(ctx, query,
		&in.TweetId,
		&in.MediaType,
		&in.Url,
		in.Id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("tweet media not found")
		}
		return err
	}

	return nil
}

func (t *tweetMedia) DeleteTweetMedia(ctx context.Context, ID string) error {
	// Query to delete the tweet permanently
	query := static.DeleteTweetMediaQuery

	result, err := t.db.Db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("tweet media not found")
	}

	return nil
}

func (t *tweetMedia) DeleteTweetMediaWithTweetID(ctx context.Context, ID string) error {
	// Query to delete the tweet permanently
	query := static.DeleteTweetMediaWithTweetIDQuery

	result, err := t.db.Db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("tweet medias not found with the tweet ID")
	}

	return nil
}
