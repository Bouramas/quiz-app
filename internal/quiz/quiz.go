package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bouramas/quiz-app/internal/question"
)

const clearScreen = "\033[2J\033[H"

func RunQuiz(questions []question.Question) {
	score := 0
	totalQuestions := len(questions)
	reader := bufio.NewReader(os.Stdin)
	quizStart := time.Now()

	for i, q := range questions {
		fmt.Print(clearScreen)
		fmt.Printf("Question %d/%d:\n%s\n", i+1, totalQuestions, q.Text)
		for j, opt := range q.Options {
			fmt.Printf("%d. %s\n", j+1, opt.Text)
		}

		fmt.Print("Enter your answer(s) (comma-separated if multiple): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userAnswers := strings.Split(input, ",")

		correct := true
		for j, opt := range q.Options {
			userSelected := contains(userAnswers, strconv.Itoa(j+1))
			if (opt.IsCorrect && !userSelected) || (!opt.IsCorrect && userSelected) {
				correct = false
				break
			}
		}

		fmt.Print(clearScreen)
		if correct {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect. The correct answer(s) was/were:")
			for j, opt := range q.Options {
				if opt.IsCorrect {
					fmt.Printf("%d. %s\n", j+1, opt.Text)
				}
			}
		}

		if i < totalQuestions-1 {
			fmt.Print("\nPress Enter to continue to the next question...")
			reader.ReadString('\n')
		} else {
			fmt.Print("\nYou answered all questions!. Press Enter to see your results.")
			reader.ReadString('\n')
		}
	}

	endTime := time.Now()
	timeSpent := endTime.Sub(quizStart)
	fmt.Print(clearScreen)
	fmt.Printf("\nQuiz completed!\nYou scored %d out of %d\nTotal time spent: %s\n", score, totalQuestions, timeSpent.Round(time.Second))
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.TrimSpace(s) == item {
			return true
		}
	}
	return false
}
