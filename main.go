package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	file, err := os.Open("internal/questions.csv")
	timeLimit := flag.Int("time", 5, "Time limit in seconds")
	flag.Parse()
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
	quizzes := createQuizzes(data)
	correctAnswers := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
problemloop:
	for _, quiz := range quizzes {
		fmt.Println(quiz.Question)
		answerChan := make(chan string)
		go func() {
			var answer string
			_, err := fmt.Scanf("%s\n", &answer)
			if err != nil {
				return
			}
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %.2f out of 100\n", (float64(correctAnswers)/float64(len(quizzes)))*100)
			break problemloop
		case answer := <-answerChan:
			if answer == quiz.Answer {
				correctAnswers++
			}
		}
	}
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
