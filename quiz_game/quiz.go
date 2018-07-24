package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 3, "t ime limit for the quiz in seconds.")
	flag.Parse()

	// open file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d - %s = \n", i+1, p.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			// write the answer into the channel
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			// if time is up, end the quiz and give results.
			fmt.Printf("You scored %d out of %d.", score, len(problems))
		case answer := <-answerChannel:
			if answer == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.", score, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // exit due to failure
}

func parseLines(lines [][]string) []problem {
	/*
		Loop through each question answer combination
		make it take the shape of the problem struct
	*/
	ret := make([]problem, len(lines)) // create a slice
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // strip out spaces
		}
	}
	return ret
}

type problem struct {
	// define question answer format
	question string
	answer   string
}
