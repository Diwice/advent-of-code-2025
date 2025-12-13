package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func is_in_ranges(rng [][]int, val int) bool {
	var in_range bool
	for i := range rng {
		if val >= rng[i][0] && val <= rng[i][1] {
			in_range = true
			break
		}
	}
	return in_range
}

func convert_ingridients(inp []string) ([]int, error) {
	var res []int
	for i := range inp {
		elem, err := strconv.Atoi(strings.TrimSpace(inp[i]))
		if err != nil {
			return []int{}, err
		}

		res = append(res, elem)
	}
	return res, nil
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

func separate_vals(inp []string) ([][]int, []int, error) {
	var rng []string
	var ings []string

	for i, v := range inp {
		if v != "\n" && strings.Contains(v, "-") {
			rng = append(rng, v)
			continue
		}
		ings = append(ings, inp[i + 1:]...)
		break
	}

	rng_out, err := convert_ranges(rng)
	if err != nil {
		return [][]int{}, []int{}, err
	}

	ings_out, err := convert_ingridients(ings)
	if err != nil {
		return [][]int{}, []int{}, err
	}

	return rng_out, ings_out, nil
}

func get_answer(inp string) (int, error) {
	splitted := strings.Split(inp, "\n")
	ranges, ingridients, err := separate_vals(splitted)
	if err != nil {
		return 0, err
	}

	var fresh_ing []int
	for _, v := range ingridients {
		if is_in_ranges(ranges, v) {
			fresh_ing = append(fresh_ing, v)
		}
	}

	return len(fresh_ing), nil
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
