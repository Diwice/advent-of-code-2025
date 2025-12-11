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

const elem_limit = 12

func find_max(inp string) (string, error) {
	var elements []string

	if len(inp) < elem_limit {
		return "", fmt.Errorf("Length of the input strings should be at least the same as element limit")
	}

	last_idx := 0
	for i := 1; i <= elem_limit; i++ {
		calc_len := len(inp) - (elem_limit - i)
		cur_max, c_max_idx, err := sub_find(inp[last_idx:calc_len])
		if err != nil {
			return "", err
		}

		last_idx += c_max_idx + 1

		elements = append(elements, cur_max)
	}

	return strings.Join(elements, ""), nil
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
