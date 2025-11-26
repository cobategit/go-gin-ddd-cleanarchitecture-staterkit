package db

import (
	"log"
	"time"

	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", cfg.Pgsql.DSN)
	if err != nil {
		log.Fatalf("cannot open postgres: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping postgres: %v", err)
	}
	return db
}
