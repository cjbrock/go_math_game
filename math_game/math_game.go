package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Question struct {
	q   string
	ans string
}

func main() {
	// read data here
	questions, err := DataIn("problems.csv")

	// log fatal errors
	if err != nil {
		log.Fatal(err)
	}

	// create each question, figure out delimiters
	for _, quests := range questions {

		quest := Question{
			q:   quests[0],
			ans: quests[1],
		}

		fmt.Printf("q: %s:, ans: %s ", quest.q, quest.ans)

	}

	// print out questions just as a test, then move on
}

func DataIn(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil

}

// parse in csv
// separate into questions
// output to user
// take in user info
// check against answer
// track correct answers
// track number of questions
// at the end of the file, output correct/total questions

// part 2:
// add timer
// default to 30 seconds
// take in optional flag
// clean white space around input
// check for non-valid input (letters, etc)
// shuffle questions
