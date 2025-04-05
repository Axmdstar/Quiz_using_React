package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"SomaliAiQuizApi/types"
)

func Live(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ok")
}

func GetQuizOptn(w http.ResponseWriter, r *http.Request) {
	diff := r.URL.Query().Get("difficulty")
	ctgy := r.URL.Query().Get("category")

	log.Printf("GetQuiz Query Params: difficulty = %v, category = %v", diff, ctgy)
	
	newQuiz, err := GetQuiz(diff, ctgy)
	if err != nil {
		responseWithERROR(w, 400, fmt.Sprintf("Failed to fetch Quiz: %v ", err))
	}
	// quizjson, err := json.Marshal(newQuiz)
	if err != nil {
		fmt.Fprint(w, err)
	}
  responsewithJSON(w, 200, newQuiz)
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
