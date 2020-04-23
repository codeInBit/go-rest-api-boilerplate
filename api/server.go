package api

import (
	"fmt"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	} else {
		fmt.Println("Success loading .env file")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seeders.Load(server.DB)
	server.Run(":", os.Getenv("SERVER_PORT"))
}
