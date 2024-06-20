package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	_ "hash/adler32"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Quiz struct {
	q, a string
}

func (q Quiz) CheckAnswer(answer string) bool {
	if answer == q.a {
		return true
	}
	return false
}

func PrintScore(correctAnswers, Questions int) {
	Score := float64(correctAnswers) / float64(Questions)
	Score *= 100
	fmt.Println("********************************")
	fmt.Printf("You Answered %d questions in total\n", Questions)
	fmt.Printf("Final Score : %.0f%% (%d/%d)", Score, correctAnswers, Questions)
}

func main() {

	var csvFile = flag.String("csv", "quiz.csv", "Provide a CSV file, Each line with this format: 'questions,answer'")
	var timeLimit = flag.Int("time", 20, "Set a time Limit for the Quiz")

	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		log.Fatal("Error while Opening CSV file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading data from the CSV file: ", err)
	}

	CorrectAnswers_count := 0
	Questions_count := 0

	Timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, line := range data {

		quiz := Quiz{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}

		fmt.Printf("Question %d: %s\n", i+1, quiz.q)
		answerChan := make(chan string)
		wg.Add(1)
		go func() {
			defer wg.Done()
			var answer string
			fmt.Print("Answer: ")
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-Timer.C:
			fmt.Printf("\n********************************")
			fmt.Print("\nTimes Up!\n")
			PrintScore(CorrectAnswers_count, Questions_count)
			return
		case answer := <-answerChan:
			if quiz.CheckAnswer(answer) {
				CorrectAnswers_count++
				Questions_count++
				fmt.Printf("Correct!\n\n")
			} else {
				Questions_count++
				fmt.Println("Incorrect!")
				fmt.Printf("The correct answer is: %s\n\n", quiz.a)
			}
		}

	}

	PrintScore(CorrectAnswers_count, Questions_count)

	wg.Wait()
}
