package db

import (
	"ipc/ent"
	"log"
	"os"
)

func EntClient() *ent.Client {

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/ipc_db?sslmode=disable"
	}

	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client
}
