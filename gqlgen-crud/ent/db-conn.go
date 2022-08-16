package ent

import (
	"context"
	"log"
)

func ConnectDB() *Client {
	client, err := Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=example password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
