package question

type Question struct {
    Text    string
    Options []Option
}

type Option struct {
    Text      string
    IsCorrect bool
}
