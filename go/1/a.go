package one

import (
	"bufio"
	"os"
	"strconv"
)

func depthsFromInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var depths []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		depths = append(depths, i)
	}

	return depths, nil
}

func A(path string) (int, error) {
	depths, err := depthsFromInput(path)
	if err != nil {
		return 0, err
	}

	return findNumIncreases(depths), nil
}

func B(path string) (int, error) {
	var (
		sums []int
	)

	depths, err := depthsFromInput(path)
	if err != nil {
		return 0, err
	}

	for i := 2; i < len(depths); i++ {
		sum := depths[i] + depths[i-1] + depths[i-2]
		sums = append(sums, sum)
	}

	return findNumIncreases(sums), nil
}

func findNumIncreases(measurements []int) int {
	var (
		numIncreases int
		lastSeen     = measurements[0]
	)

	for i := 1; i < len(measurements); i++ {
		measurement := measurements[i]
		if measurement > lastSeen {
			numIncreases++
		}
		lastSeen = measurement
	}
	return numIncreases
}
