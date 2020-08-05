package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// We define a struct for the questions

type questionAnswer struct {
	Question string
	Answer   string // This usually is a number, since it is a math quiz. However, we stick to string to keep it more general.
}

func main() {

	// Get commandline arguments
	csvFilename := flag.String("quizfile", "problems.csv", "CSV file containing the questions")

	flag.Parse()

	// Open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		// If there is an error we log it
		log.Fatal(err)
	}

	// We read the file as csv
	csvReader := csv.NewReader(bufio.NewReader(file))

	// Score of the user
	score := 0

	// Total questions
	total_questions := 0

	// Input reader for the user's answer
	stdinReader := bufio.NewReader(os.Stdin)

	// We go through each entry in csv file
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

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

		total_questions++
	}

	fmt.Println("You scored", score, "points! There were", total_questions, "questions in total!")

}
