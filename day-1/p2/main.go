package main

import (
	"os"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part2(text string) string {
	dial := 50
	password := 0
	splitted := strings.Split(text, "\n")
	for _, text := range splitted {
		direction := text[0]
		value, _ := strconv.Atoi(strings.TrimSpace(text[1:]))

		zeroesPassed := 0

		if direction == 'R' {
			zeroesPassed = int(math.Abs(float64(dial+value)) / 100)
			dial = ((dial + value) % 100 + 100) % 100
		} else {
			inverted := 0
			if dial > 0 {
				inverted = (100 - dial)
			}
			zeroesPassed = int(math.Abs(float64(inverted+value)) / 100)
			dial = ((dial - value) % 100 + 100) % 100
		}

		password += zeroesPassed
	}

	return fmt.Sprint(password)
}

func main() {
	inp, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Couldn't read the input: %v\n", err)
		return
	}

	fmt.Printf("Answer is: %v\n", part2(string(inp)))
}
