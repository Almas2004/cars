package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

func Load() Config {
	_ = godotenv.Load()
	db := os.Getenv("DATABASE_URL")

	if db == "" {
		log.Fatalf("Url is empty")
	}

	return Config{
		DBUrl: db,
	}
}