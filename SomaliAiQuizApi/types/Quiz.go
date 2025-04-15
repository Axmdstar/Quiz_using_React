package types

import (
	"errors"
)

type Quiz struct {
	ResponseCode int      `json:"response_code"`
	Results      []Result `json:"results"`
}

type Result struct {
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Category         string   `json:"category"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func NewQuiz() *Quiz {
	return &Quiz{}
}

type ResponseQuiz struct {
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Category         string   `json:"category"`
	Question         Question `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type Question struct {
	Som string `json:"som"`
	Eng string `json:"eng"`
}

func NEWResponseQuiz() *ResponseQuiz {
	return &ResponseQuiz{}
}

const TRANSLATESTRING = `translate the Questions to somali, don't change the structure of the json only the question, return in this Format
  {
  "response_code": 0,
  "results": [
    {
      "type": "",
      "difficulty": "",
      "category": "",
      "question": "",
      "correct_answer": "",
      "incorrect_answers": []
    },
   ]
  }
  `

// this func Merges the (q)Generated somali quiz and (d)default quiz to
// get a new stuct with somali and english Question
func MergeQuestions(somQue Result, engQue Result) (*ResponseQuiz, error) {
	if somQue.Question == "" || engQue.Question == "" {
		return nil, errors.New(" Quiz Not Found")
	}

	resQuiz := NEWResponseQuiz()

	resQuiz.Category = somQue.Category
	resQuiz.CorrectAnswer = somQue.CorrectAnswer
	resQuiz.Difficulty = somQue.Difficulty
	resQuiz.IncorrectAnswers = somQue.IncorrectAnswers
	resQuiz.Type = somQue.Type

	resQuiz.Question.Som = somQue.Question
	resQuiz.Question.Eng = engQue.Question

	return resQuiz, nil
}
