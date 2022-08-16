package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"odin/ent"
	"os"
)

var dns string

func init() {
	dns = fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
	)
}

func Conn() (*ent.Client, error) {
	client, err := ent.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	client = client.Debug()

	return client, nil
}
