package db

import (
	"log"
	"time"

	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLDB(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Open("mysql", cfg.Mysql.DSN)
	if err != nil {
		log.Fatalf("cannot open mysql: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping mysql: %v", err)
	}
	return db
}
