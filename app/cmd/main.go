package main

import (
	"fmt"
	"log"
	"net/http"
	"vcard/Database"
	router "vcard/api"

	"github.com/joho/godotenv"
)

func main() {
	initDotenv()
	Database.InitDatabaseCredentials()
	fmt.Println("Server is running on port 8081")
	http.HandleFunc("/", router.Home)
	http.ListenAndServe(":8081", nil)
}

func initDotenv() {
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Print("No .env file found")
		} else {
			log.Print("Using .env file found in parent dir")
		}
	} else {
		log.Println("Using .env file found in current dir")
	}
}
