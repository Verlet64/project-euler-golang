package problem67maximumpathsumtwo

import (
	"strconv"
	"strings"
)

func (s *Solver) solve () (int, error) {
	parsed, err := s.parse(s.input)
	if err != nil {
		return -1, err
	}

	rowCount := len(parsed)

	for i := len(parsed[rowCount - 2]) - 1; i >= 0; i-- {
		cur := i
		next := i + 1

		for j := 1; j < len(parsed[next]); j++ {
			m := j - 1
			n := j

			if parsed[next][m] > parsed[next][n] {
				parsed[cur][m] = parsed[cur][m] + parsed[next][m]
			} else {
				parsed[cur][m] = parsed[cur][m] + parsed[next][n]
			}
		}
	}

	return parsed[0][0], nil
}

// Solver takes an input and solves problem 67
type Solver struct {
	input string
}

// New initialises a solver
func New (input string) Solver {
	return Solver{input: input}
}

func (s *Solver) parse (input string) ([][]int, error) {
	var parsed [][]int;

	rows := strings.Split(input, "\n")
	parsed = make([][]int, len(rows))

	for i, row := range(rows) {
		sanitised := strings.Trim(row, "\n")
		tokenised := strings.Split(sanitised, " ")

		if (len(tokenised) == 0) {
			continue;
		}

		parsed[i] = make([]int, len(tokenised))
		for j, element := range(tokenised) {
			num, err := strconv.Atoi(element)
			if err != nil {
				return nil, err;
			}


			parsed[i][j] = num
		}
	}

	return parsed, nil
}

