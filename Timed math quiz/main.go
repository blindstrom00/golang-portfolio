package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFile := flag.String("csv", "questions.csv", "CSV File formated as question,answer")
	flag.Parse()
	timeLimit := flag.Int("limit", 60, "The time limit in seconds")

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFile))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to open the CSV file provided")
	}

	problems := lineParser(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	fmt.Printf("You have %d Seconds!\n", *timeLimit)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime is up!\nYou got %d out of %d correct!\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}
	fmt.Printf("You got %d out of %d correct!\n", correct, len(problems))
}

func lineParser(lines [][]string) []problem {
	rturn := make([]problem, len(lines))
	for i, line := range lines {
		rturn[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return rturn
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
