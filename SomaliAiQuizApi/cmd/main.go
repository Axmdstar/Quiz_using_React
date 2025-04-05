package main

import (
	"log"
	"net/http"
	"os"

	"SomaliAiQuizApi/api"

	"github.com/joho/godotenv"
)

func main() {
	// Get env File and Variables in it
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("ENV NOT FOUND")
	}

	PORT := os.Getenv("PORT")

	// start Our router multiplexer for handling http request endpoint 
	router := http.NewServeMux()
	
	// Routes
	router.HandleFunc("/getquiz", api.GetQuizOptn)

	// Server Option
	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	// Start Server
	log.Printf("Server Listining On Port: %v", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to Start Server : %v ", err)
	}
}
