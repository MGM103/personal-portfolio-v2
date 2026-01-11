package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mgm103/personal-portfolio-v2/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	server := app.NewServer(port)

	fmt.Println("Server running on PORT:", port)
	log.Fatal(server.ListenAndServe())
}
