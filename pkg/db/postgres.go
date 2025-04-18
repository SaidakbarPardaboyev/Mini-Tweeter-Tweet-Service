// Package postgres implements postgres connection.
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/SaidakbarPardaboyev/Mini-Tweeter-Users-Service/config"
)

// Postgres -.
type Postgres struct {
	connAttempts int
	connTimeout  time.Duration
	Builder      squirrel.StatementBuilderType
	Db           *pgxpool.Pool
}

// New -.
// func New(cfg config.Config, opts ...Option) (*Postgres, error) {
// 	pg := &Postgres{
// 		connAttempts: _defaultConnAttempts,
// 		connTimeout:  _defaultConnTimeout,
// 	}

// 	// Custom options
// 	for _, opt := range opts {
// 		opt(pg)
// 	}

// 	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
// 		cfg.PostgresUser,
// 		cfg.PostgresPassword,
// 		cfg.PostgresHost,
// 		cfg.PostgresPort,
// 		cfg.PostgresDatabase,
// 	)

// 	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
// 	var err error
// 	for pg.connAttempts > 0 {
// 		pg.Db, err = sqlx.Connect("postgres", pgxUrl)
// 		if err == nil {
// 			break
// 		}

// 		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

// 		time.Sleep(pg.connTimeout)

// 		pg.connAttempts--
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
// 	}

// 	return pg, nil
// }

// Close -.
// func (p *Postgres) Close() {
// 	if p.Db != nil {
// 		p.Db.Close()
// 	}
// }

func NewPostgresDB(cfg *config.Config) (*Postgres, error) {
	response := &Postgres{}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Optional: Customize pool configuration
	config.MaxConns = 1000
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	response.Db = pool
	response.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return response, nil
}
