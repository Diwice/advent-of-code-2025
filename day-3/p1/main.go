package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func sub_find(inp string) (string, int, error) {
	max, max_idx := 0, 0
	for i, v := range inp {
		int_val, err := strconv.Atoi(string(v))
		if err != nil {
			return "", 0, err
		}

		if int_val > max {
			max = int_val
			max_idx = i
		}
	}

	conv_max := strconv.Itoa(max)

	return conv_max, max_idx, nil
}

func find_max(inp string) (string, error) {
	f_max, f_max_idx, err := sub_find(inp)
	if err != nil {
		return "", err
	}

	if f_max_idx == len(inp) - 1 {
		f_max, f_max_idx, err = sub_find(inp[:len(inp) - 1])
		if err != nil {
			return "", err
		}
	}

	s_max, _, err := sub_find(inp[f_max_idx + 1:])
	if err != nil {
		return "", err
	}

	return f_max + s_max, nil
}

func sum(inp []int) int {
	res := 0
	for i := range inp {
		res += inp[i]
	}
	return res
}

func get_answer(inp string) (int, error) {
	var joltages []int
	lines := strings.Split(inp, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		max, err := find_max(line)
		if err != nil {
			return 0, err
		}

		in_int, err := strconv.Atoi(max)
		if err != nil {
			return 0, err
		}

		joltages = append(joltages, in_int)
	}

	return sum(joltages), nil
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Couldn't open the file:", err)
		return
	}

	answer, err := get_answer(string(file))
	if err != nil {
		fmt.Println("Couldn't get the answer:", err)
		return
	}

	fmt.Println("The answer is:", answer)
}
