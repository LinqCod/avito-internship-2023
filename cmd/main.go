package main

import (
	"github.com/linqcod/avito-internship-2023/pkg/config"
	"github.com/linqcod/avito-internship-2023/pkg/database"
	"log"
)

func init() {
	config.LoadConfig(".env")
}

func main() {
	// init db connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("error while trying to ping db: %v", err)
	}

	log.Println("Success ping!")

	// TODO: init routing

	// TODO: init swagger

	// TODO: add repos, services, workers, etc

	// TODO: init server

	// TODO: graceful shutdown
}
