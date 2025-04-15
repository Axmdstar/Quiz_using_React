package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"SomaliAiQuizApi/types"

	"github.com/google/generative-ai-go/genai"
)

func Live(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ok")
}

func GetQuizOptn(ctx context.Context, gm *genai.GenerativeModel) http.HandlerFunc {
	// Return Http Handler function, main function has access to passed ctx, gm params
	return func(w http.ResponseWriter, r *http.Request) {
		// get difficlty and category query params
		diff := r.URL.Query().Get("difficulty")
		ctgy := r.URL.Query().Get("category")

		log.Printf("GetQuiz Query Params: difficulty = %v, category = %v", diff, ctgy)

		// get Quiz from Opentdb api
		newQuiz, err := GetQuiz(diff, ctgy)
		if err != nil {
			responseWithERROR(w, 400, fmt.Sprintf("Failed to fetch Quiz: %v ", err))
		}

		// join Translate string and quiz json
		prptString := fmt.Sprintf(`%v %v`, types.TRANSLATESTRING, newQuiz)
		// Create content with Gemini
		gm.ResponseMIMEType = "application/json"
		resp, err := gm.GenerateContent(ctx, genai.Text(prptString))
		if err != nil {
			responseWithERROR(w, 404, fmt.Sprint(err))
		}

		joinpart := JoinParts(*resp.Candidates[0].Content)
		generatedQuiz := types.NewQuiz()

		log.Printf("Parts Joined >> %v ", joinpart)
		err = json.Unmarshal([]byte(joinpart), generatedQuiz)
		if err != nil {
			fmt.Printf("Error >> %v", err)
		}
		// fmt.Printf("\n >>> %v \n ", generatedQuiz)
		var resQuiz []*types.ResponseQuiz
		for i, q := range newQuiz.Results {
			nq, err := types.MergeQuestions(q, generatedQuiz.Results[i])
			if err != nil {
				log.Println(err)
			}
			resQuiz = append(resQuiz, nq)
		}

		responsewithJSON(w, 200, resQuiz)
	}
}

func JoinParts(p genai.Content) string {
	var joinVar string
	for _, parts := range p.Parts {
		joinVar += fmt.Sprint(parts)
	}
	return fmt.Sprintf(`%v`, joinVar)
}

func GetQuiz(diff string, ctgy string) (*types.Quiz, error) {
	dbQuiz := "https://opentdb.com/api.php?amount=10"
	if diff != "" {
		dbQuiz += fmt.Sprintf("&difficulty=%v", diff)
	}
	if ctgy != "" {
		dbQuiz += fmt.Sprintf("&category=%v", ctgy)
	}

	fmt.Println(dbQuiz)
	resp, err := http.Get(dbQuiz)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	readdata, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	quiz := types.NewQuiz()
	err = json.Unmarshal(readdata, quiz)
	if err != err {
		return nil, err
	}

	return quiz, nil
}
