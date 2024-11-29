package main

import (
	"fmt"

	"github.com/bouramas/quiz-app/internal/question"
	"github.com/bouramas/quiz-app/internal/quiz"
)

func main() {
	questions, err := question.LoadQuestions("questions.txt")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		return
	}

	quiz.RunQuiz(questions)
}
