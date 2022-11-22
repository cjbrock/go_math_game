package mathgame

import (
	"encoding/csv"
	"os"
)

type Question struct {
	q1  int
	q2  int
	ans int
}

func main() {
	// read data here
	// log fatal errors
	// create each question, figure out delimiters
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
