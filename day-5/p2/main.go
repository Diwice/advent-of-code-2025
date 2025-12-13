package main

import (
	"os"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func sorting_func(rng_one, rng_two []int) int {
	if rng_one[0] > rng_two[0] {
		return 1
	} else if rng_one[0] < rng_two[0] {
		return -1
	}

	if rng_one[1] > rng_two[1] {
		return 1
	} else if rng_one[1] < rng_two[1] {
		return -1
	}

	return 0
}

func count_unique_ids(inp [][]int) int {
	slices.SortFunc(inp, sorting_func)

	rng_copy := make([][]int, len(inp))
	for i := range inp {
		rng_copy[i] = make([]int, 2)
		copy(rng_copy[i], inp[i])
	}

	merged_ranges := make([][]int, 0)
	cur_elem := make([]int, 2)
	copy(cur_elem, rng_copy[0])

	for i := 1; i < len(rng_copy); i++ {
		next_elem := rng_copy[i]

		if next_elem[0] <= cur_elem[0] + 1 {
			if next_elem[1] > cur_elem[1] {
				cur_elem[1] = next_elem[1]
			}
		} else {
			merged_ranges = append(merged_ranges, cur_elem)
			cur_elem = make([]int, 2)
			copy(cur_elem, next_elem)
		}
	}

	merged_ranges = append(merged_ranges, cur_elem)

	var unique_ids int
	for _, v := range merged_ranges {
		rng_len := v[1] - v[0] + 1
		unique_ids += rng_len
	}

	return unique_ids
}

func convert_ranges(inp []string) ([][]int, error) {
	var res [][]int
	for i := range inp {
		splitted := strings.Split(strings.TrimSpace(inp[i]), "-")

		elem_one, err := strconv.Atoi(splitted[0])
		if err != nil {
			return [][]int{}, err
		}

		elem_two, err := strconv.Atoi(splitted[1])
		if err != nil {
			return [][]int{}, err
		}

		res = append(res, []int{elem_one, elem_two})
	}
	return res, nil
}

func separate_vals(inp []string) ([][]int, error) {
	var rng []string

	for _, v := range inp {
		if v != "\n" && strings.Contains(v, "-") {
			rng = append(rng, v)
			continue
		}
		break
	}

	rng_out, err := convert_ranges(rng)
	if err != nil {
		return [][]int{}, err
	}

	return rng_out, nil
}

func get_answer(inp string) (int, error) {
	splitted := strings.Split(inp, "\n")
	ranges, err := separate_vals(splitted)
	if err != nil {
		return 0, err
	}

	total_fresh := count_unique_ids(ranges) // ngl sounds like some toothpaste...

	return total_fresh, nil
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Couldn't read file:", err)
		return
	}

	answer, err := get_answer(string(file))
	if err != nil {
		fmt.Println("Couldn't retrieve answer:", err)
		return
	}

	fmt.Println("The answer is:", answer)
}
