package three

import (
	"fmt"
	"strconv"

	"github.com/isaachess/advent2021/common"
)

func A(path string) (int64, error) {
	diagnostics, err := common.FileData(path)
	if err != nil {
		return 0, err
	}

	numPerColumn := make([]int, len(diagnostics[0]))

	for _, diagnostic := range diagnostics {
		for i, c := range diagnostic {
			if c == '1' {
				numPerColumn[i]++
			}
		}
	}

	var (
		gammaRate   string
		epsilonRate string
	)

	comparison := len(diagnostics) / 2

	for i, num := range numPerColumn {
		// Instructions do not mention this possibility. So sanity-check that it
		// does not occur.
		if num == comparison {
			return 0, fmt.Errorf("bit value exactly equal for column %d", i)
		}

		if num > comparison {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		return 0, err
	}

	return gamma * epsilon, nil
}

func B(path string) (int, error) {
	diagnostics, err := common.FileData(path)
	if err != nil {
		return 0, err
	}

	numPerColumn := make([]int, len(diagnostics[0]))

	for _, diagnostic := range diagnostics {
		for i, c := range diagnostic {
			if c == '1' {
				numPerColumn[i]++
			}
		}
	}

	oxygenGeneratorRating, err := filter(diagnostics, 0, true)
	if err != nil {
		return 0, err
	}
	co2ScrubberRating, err := filter(diagnostics, 0, false)
	if err != nil {
		return 0, err
	}

	return oxygenGeneratorRating * co2ScrubberRating, nil

}

func filter(diagnostics []string, column int, wantMajority bool) (int, error) {
	numPerColumn := make([]int, len(diagnostics[0]))

	for _, diagnostic := range diagnostics {
		for i, c := range diagnostic {
			if c == '1' {
				numPerColumn[i]++
			}
		}
	}

	comparison := len(diagnostics) / 2

	// If it's an odd number, the divide by two rounds down so we actually have
	// wrong number for comparison. So add one.
	if len(diagnostics)%2 != 0 {
		comparison++
	}

	num := numPerColumn[column]
	oneIsMajority := num >= comparison
	var filtered []string
	for _, diagnostic := range diagnostics {
		var hasMajority bool
		if oneIsMajority {
			hasMajority = diagnostic[column] == '1'
		} else {
			hasMajority = diagnostic[column] == '0'
		}
		if hasMajority == wantMajority {
			filtered = append(filtered, diagnostic)
		}
	}
	if len(filtered) == 1 {
		converted, err := strconv.ParseInt(filtered[0], 2, 64)
		return int(converted), err
	}
	return filter(filtered, column+1, wantMajority)
}
