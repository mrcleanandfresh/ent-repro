package main

import (
	"fmt"
	"log"
	"test/db"
	_ "test/db/runtime"
)

func main() {
	database, err1 := db.ConnectToDb()
	defer func(database *db.Client) {
		err := database.Close()
		if err != nil {
			log.Fatal(fmt.Errorf("could not close the database connection: %w", err))
		}
	}(database)

	if err1 != nil {
		log.Fatal(err1)
		return
	}
}
