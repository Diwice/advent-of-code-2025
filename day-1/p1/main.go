package main

import (
	"os"
	"fmt"
	"strings"
)

func conv_to_num(input string) int {
	out := 0
	for _, v := range input {
		out = out*10 + int(v - '0')
	}
	return out
}

func build_num_from_str(input string) string {
	out := ""
	for _, v := range input {
		if v >= 48 && v <= 57 {
			out += string(v)
		}
	}
	return out
}

func normalize_count(input, rc *int) {
	for *input > 99 || *input < 0 {
		if *input < 0 {
			*input += 100
			continue
		}
		*input -= 100
	}

	if *input == 0 {
		*rc += 1
	}
}

func calc_from_input(input string) int {
	res := 0
	dial_count := 50
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		num_str := build_num_from_str(v)
		if v[0] == 'L' {
			dial_count -= conv_to_num(num_str)
			continue
		}

		dial_count += conv_to_num(num_str)

		normalize_count(&dial_count, &res)
	}

	return res
}

func main() {
	inp, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Couldn't read the input: %v\n", err)
		return
	}

	fmt.Printf("Answer is: %v\n", calc_from_input(string(inp)))
}
