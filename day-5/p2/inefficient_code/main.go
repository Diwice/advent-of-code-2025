package main

import (
	"os"
	"fmt"
	"sync"
	"strconv"
	"strings"
)

func work_on_map(inp map[int]bool, rng []int, mu *sync.Mutex) {
	for j := rng[0]; j <= rng[1]; j++ {
		mu.Lock()
		inp[j] = true
		mu.Unlock()
	}
}

func fill_and_count_ids(inp [][]int) int {
	n_map := map[int]bool{}

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := range inp {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			work_on_map(n_map, inp[idx], &mu)
		}(i)
	}

	wg.Wait()

	return len(n_map)
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

	fresh_ings := fill_and_count_ids(ranges)

	return fresh_ings, nil
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
