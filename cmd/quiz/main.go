package main

import (
	"fmt"
	"os"

	"github.com/bouramas/quiz-app/internal/question"
	"github.com/bouramas/quiz-app/internal/quiz"
)

func main() {

	filename := os.Args[1]
	questions, err := question.LoadQuestions(filename)
	if err != nil {
		fmt.Println("Error loading questions:", err)
		return
	}

	quiz.RunQuiz(questions)
}
