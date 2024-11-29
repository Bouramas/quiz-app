package question

import (
    "bufio"
    "os"
    "strings"
)

func LoadQuestions(filename string) ([]Question, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var questions []Question
    var currentQuestion Question
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            if currentQuestion.Text != "" {
                questions = append(questions, currentQuestion)
                currentQuestion = Question{}
            }
        } else if currentQuestion.Text == "" {
            currentQuestion.Text = line
        } else {
            isCorrect := strings.HasSuffix(line, "*")
            text := strings.TrimSuffix(line, "*")
            currentQuestion.Options = append(currentQuestion.Options, Option{Text: text, IsCorrect: isCorrect})
        }
    }

    if currentQuestion.Text != "" {
        questions = append(questions, currentQuestion)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return questions, nil
}
