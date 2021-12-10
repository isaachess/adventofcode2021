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
