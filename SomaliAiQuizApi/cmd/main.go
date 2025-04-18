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

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins (replace * with specific domains in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		// Allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Get env File and Variables in it
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("ENV NOT FOUND")
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

	handler := enableCORS(router)
	// Server Option
	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: handler,
	}

	// Start Server
	log.Printf("Server Listining On Port: %v", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to Start Server : %v ", err)
	}
}
