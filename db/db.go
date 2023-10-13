package db

import (
	"context"
	"fmt"
	"log"
	"os"
	//_ "test/db/runtime"
)

var host = os.Getenv("DB_HOST")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var table = os.Getenv("DB_TABLE")
var port = os.Getenv("DB_PORT")

func ConnectToDb() (*Client, error) {
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, table, port)
	client, err := Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	return client, nil
}
