package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// We define a struct for the questions

type questionAnswer struct {
	Question string
	Answer   string // This usually is a number, since it is a math quiz. However, we stick to string to keep it more general.
}

func showScoreAndExit(score int, totalQuestions int) {
	fmt.Println("You scored", score, "points! There were", totalQuestions, "questions in total!")
	os.Exit(0)
}

func main() {

	// Get commandline arguments
	csvFilename := flag.String("quizfile", "problems.csv", "CSV file containing the questions")
	timeout := flag.Int("timeout", 30, "Maximum time to answer all questions")

	flag.Parse()

	// Open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		// If there is an error we log it
		log.Fatal(err)
	}

	// We read the file as csv
	csvReader := csv.NewReader(bufio.NewReader(file))

	// We read in all questions
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Score of the user
	score := 0

	// Total questions
	totalQuestions := len(records)

	// Input reader for the user's answer
	stdinReader := bufio.NewReader(os.Stdin)

	// Tell the user the game is about to start
	fmt.Println("The game starts now! You only have", *timeout, "seconds to answer all question. Once you are ready hit the ENTER key!")
	fmt.Scanln()

	// We create a timer
	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	// We create a function which shows the score. This function is run in parallel, so the questions are still shown.
	go func() {
		<-timer.C
		showScoreAndExit(score, totalQuestions)
	}()

	// We go through each entry in csv file
	for _, record := range records {

		// We put the record into a QuestionAnswer element
		questionAnswer := questionAnswer{
			Question: record[0],
			Answer:   record[1],
		}

		// Let's ask the user for the answer for this question
		fmt.Println(questionAnswer.Question)

		// Read in the answer
		useranswer, _ := stdinReader.ReadString('\n')
		if strings.Contains(useranswer, "\r") {
			useranswer = strings.Replace(useranswer, "\r\n", "", -1)
		}

		if useranswer == questionAnswer.Answer {
			score++
		}

	}
	showScoreAndExit(score, totalQuestions)
}
