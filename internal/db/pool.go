package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(databaseUrl string) *pgxpool.Pool {
	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.New(c, databaseUrl)
	if err != nil {
		log.Fatalf("pgxpool.New error: %v", err)
	}
	if err = pool.Ping(c); err != nil {
		log.Fatalf("Ping is error %v", err)
	}

	return pool
}