package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

func commandsFromFile(path string) ([]command, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var commands []command
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			return nil, fmt.Errorf("expected line length 2, got %d", len(line))
		}
		iVal, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		commands = append(commands, command{
			direction: line[0],
			value:     iVal,
		})
	}

	return commands, nil
}

func A(path string) (int, error) {
	commands, err := commandsFromFile(path)
	if err != nil {
		return 0, err
	}

	hPos := 0
	depth := 0
	for i, comm := range commands {
		switch comm.direction {
		case "forward":
			hPos += comm.value
		case "down":
			depth += comm.value
		case "up":
			depth -= comm.value
		}
		// Sanity check
		if hPos < 0 || depth < 0 {
			return 0, fmt.Errorf("negative pos reached at step %d", i)
		}
	}
	return hPos * depth, nil
}

func B(path string) (int, error) {
	commands, err := commandsFromFile(path)
	if err != nil {
		return 0, err
	}

	hPos := 0
	aim := 0
	depth := 0
	for i, comm := range commands {
		switch comm.direction {
		case "forward":
			hPos += comm.value
			depth += comm.value * aim
		case "down":
			aim += comm.value
		case "up":
			aim -= comm.value
		}
		// Sanity check
		if hPos < 0 || depth < 0 || aim < 0 {
			return 0, fmt.Errorf("negative val reached at step %d", i)
		}
	}
	return hPos * depth, nil
}
