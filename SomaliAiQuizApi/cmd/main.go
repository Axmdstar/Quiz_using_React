package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"SomaliAiQuizApi/api"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	// Get env File and Variables in it
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("ENV NOT FOUND")
	}

	PORT := os.Getenv("PORT")
	APIKEY := os.Getenv("APIKEY")

	// start Our router multiplexer for handling http request endpoint
	router := http.NewServeMux()

	startgemini := api.GeminiAPI{Model: "gemini-2.0-flash", APIKEY: APIKEY}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(APIKEY))
	if err != nil {
		log.Fatalln(err)
	}
	// Close Client
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	// Routes
	router.HandleFunc("/getquiz", api.GetQuizOptn(ctx, model))
	router.HandleFunc("/dummy", api.DummyData)
	router.HandleFunc("/testai", startgemini.GeminiMiddleware())

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
