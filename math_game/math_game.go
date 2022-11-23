package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)



func main() {
	// read data here
	questions, err := DataIn("problems.csv")
	// log errors
	if err != nil {
		quit("Failed to parse the provided CSV file.")
	}
	problems := parseRows(questions)

}

func parseRows(rows [][]string) []question {
	fmt.Printf("%s", rows)
	// create each question
	// print out questions just as a test, then move on
	askme := make([]question, len(rows)) 
		for i, row := range rows {
			askme[i] := question{
				q:   row[0],
				ans: row[1],
			}
			fmt.Printf("q: %s:, ans: %s ", askme.q, askme.ans)
		}
		// call a function that outputs the questions, checks the answers, and then outputs the
		
	return askme
}

type question struct {
	q   string
	ans string
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

func quit(mess string) {
	fmt.Println(mess)
	os.Exit(1)
}

// parse in csv - DONE
// separate into questions - DONE
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
