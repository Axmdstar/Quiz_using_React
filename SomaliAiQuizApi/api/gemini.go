package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiAPI struct {
	APIKEY string
	Model  string
}

func (gm *GeminiAPI) Prompt(prptString string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(gm.APIKEY))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	model := client.GenerativeModel(gm.Model)
	resp, err := model.GenerateContent(ctx, genai.Text(prptString))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func (gm *GeminiAPI) GeminiMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		translatePrompt := "translate the content to somali, no need for explaining"
		join := fmt.Sprintf(`%v {
      "type": "multiple",
      "difficulty": "hard",
      "category": "Entertainment: Film",
      "question": "In what Disney movie can you spot the character &quot;Pac-Man&quot; in if you look closely enough in some scenes?",
      "correct_answer": "Tron",
      "incorrect_answers": [
        "Big Hero 6",
        "Fantasia",
        "Monsters, Inc."
      ]
    },`, translatePrompt)
		gm.Prompt(join)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
