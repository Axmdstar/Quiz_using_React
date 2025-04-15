package api_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"SomaliAiQuizApi/types"
)

func assertFieldCheck(t testing.TB, got, want types.Quiz) {
	t.Helper()
	// call merge function

	for i, q := range got.Results {
		res, err := types.MergeQuestions(q, want.Results[i])
		if err != nil {
			fmt.Printf("Error on index:  %d,  %v", i, q)
		}
		fmt.Printf("Merged \nEnglish ::> %v \nSomali ::> %v\n", res.Question.Eng, res.Question.Som)
	}
}

func TestQuizApi(t *testing.T) {
	fmt.Println("Reading Json File")

	jsonData, err := os.Open("../data.json")
	if err != nil {
		t.Fatal(err)
	}

	defaultData, err := os.Open("../default.json")
	if err != nil {
		t.Fatal(err)
	}

	defer defaultData.Close()
	defer jsonData.Close()

	readjson, err := io.ReadAll(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	readdefaultjson, err := io.ReadAll(defaultData)
	if err != nil {
		t.Fatal(err)
	}

	defaultDataStruct := types.NewQuiz()
	quizstruc := types.NewQuiz()

	err = json.Unmarshal(readdefaultjson, defaultDataStruct)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(readjson, quizstruc)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Som & Eng Question ", func(t *testing.T) {
		assertFieldCheck(t, *quizstruc, *defaultDataStruct)
	})
}
