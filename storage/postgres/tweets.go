package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
	pb "github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/genproto/tweets_service"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/db"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/logger"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/pkg/static"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/storage/repo"
	"github.com/google/uuid"
)

type tweets struct {
	db  *db.Postgres
	cff *config.Config
	log logger.Logger
}

func NewTweetsRepo(db *db.Postgres, log logger.Logger, cff *config.Config) repo.TweetsRepo {
	return &tweets{db: db, log: log, cff: cff}
}

func (t *tweets) CreateTweet(ctx context.Context, in *pb.Tweet) (string, error) {
	var (
		id = uuid.New().String()

		// Insert the new user into the database
		insertQuery = static.InsertTweetQuery
	)

	_, err := t.db.Db.Exec(ctx, insertQuery,
		id,
		in.Content,
		in.UserId,
	)
	if err != nil {
		return "", fmt.Errorf("error creating tweet: %v", err)
	}

	return id, nil
}

func (t *tweets) GetTweet(ctx context.Context, ID string) (*pb.Tweet, error) {
	query := static.GetTweetByIDQuery

	var (
		tweet                pb.Tweet
		createdAt, updatedAt sql.NullTime
	)

	err := t.db.Db.QueryRow(ctx, query, ID).Scan(
		&tweet.Id,
		&tweet.Content,
		&tweet.UserId,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tweet not found")
		}
		return nil, fmt.Errorf("error getting tweet: %v", err)
	}

	// Formatting timestamps
	if createdAt.Valid {
		tweet.CreatedAt = createdAt.Time.Format(time.RFC3339)
	}
	if updatedAt.Valid {
		tweet.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &tweet, nil
}

func (t *tweets) GetListTweet(ctx context.Context, in *pb.GetListTweetRequest) (*pb.TweetList, error) {
	var (
		whereClauses []string
		args         []interface{}
		argIndex     = 1

		queryBase  = static.GetListTweetBaseQuery
		queryGroup = static.GetListTweetGroupQuery
		queryCount = static.GetListTweetQueryCount
	)

	// Full-text search on description, title, fullname, and surname
	if in.Search != nil {
		whereClauses = append(whereClauses, fmt.Sprintf(
			" (content ILIKE $%d)",
			argIndex,
		))
		args = append(args, "%"+*in.Search+"%")
		argIndex += 1
	}

	// Append WHERE clause if filters exist
	if len(whereClauses) > 0 {
		queryBase += " WHERE " + strings.Join(whereClauses, " AND ")
		queryCount += " WHERE " + strings.Join(whereClauses, " AND ")
	}
	queryBase += queryGroup

	// Sorting
	sortBy := "created_at"
	if in.SortBy != nil {
		sortBy = *in.SortBy
	}
	order := "DESC"
	if in.Order != nil && (*in.Order == "asc" || *in.Order == "ASC") {
		order = "ASC"
	}
	queryBase += fmt.Sprintf(" ORDER BY %s %s", sortBy, order)

	// Pagination
	queryBase += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, in.Limit, (in.Page-1)*in.Limit)

	// Count
	var count int64
	err := t.db.Db.QueryRow(ctx, queryCount, args[:len(args)-2]...).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error fetching tweet count: %v", err)
	}

	rows, err := t.db.Db.Query(ctx, queryBase, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching tweet list: %v", err)
	}
	defer rows.Close()

	var tweets []*pb.Tweet
	for rows.Next() {
		var (
			tweet                pb.Tweet
			createdAt, updatedAt sql.NullTime
		)

		err := rows.Scan(
			&tweet.Id,
			&tweet.Content,
			&tweet.UserId,
			&createdAt,
			&updatedAt,
			&tweet.Medias,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning tweet: %v", err)
		}

		if createdAt.Valid {
			tweet.CreatedAt = createdAt.Time.Format(time.RFC3339)
		}
		if updatedAt.Valid {
			tweet.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
		}

		tweets = append(tweets, &tweet)
	}

	// Check for row iteration errors
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tweeta: %v", err)
	}

	return &pb.TweetList{
		Tweets: tweets,
		Count:  count,
	}, nil
}

func (t *tweets) UpdateTweet(ctx context.Context, in *pb.Tweet) error {
	// Query to update tweet details
	query := static.UpdateTweetQuery

	_, err := t.db.Db.Exec(ctx, query,
		&in.Content,
		&in.UserId,
		in.Id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("tweet not found")
		}
		return err
	}

	return nil
}

func (t *tweets) DeleteTweet(ctx context.Context, ID string) error {
	// Query to delete the tweet permanently
	query := static.DeleteTweetQuery

	result, err := t.db.Db.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("tweet not found")
	}

	return nil
}
