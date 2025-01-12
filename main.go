package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	file, err := os.Open("internal/questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(os.Stdin)
	quizzes := createQuizzes(data)
	correctAnswers := 0
	for _, quiz := range quizzes {
		fmt.Println(quiz.Question)
		input.Scan()
		if input.Text() == quiz.Answer {
			correctAnswers++
		}
	}
	score := (float64(correctAnswers) / float64(len(quizzes))) * 100
	fmt.Printf("%.2f\n", score)
}

func createQuizzes(data [][]string) []Quiz {
	var quizzes []Quiz
	for _, question := range data {
		var quiz Quiz
		for i, field := range question {
			if i == 0 {
				quiz.Question = field
			} else if i == 1 {
				quiz.Answer = field
			}
		}
		quizzes = append(quizzes, quiz)
	}
	return quizzes
}
