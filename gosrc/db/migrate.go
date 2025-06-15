package db

import (
	"context"
	"log"

	_ "github.com/lib/pq"
)

func MigrateDatabase() {
	client := EntClient()
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		panic("failed creating schema resources: " + err.Error())
	}
	log.Println("Database migration completed successfully.")

}
