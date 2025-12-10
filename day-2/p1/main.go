package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func are_equal(inp []string) bool {
	initial := inp[0]
	for i := 1; i < len(inp); i++ {
		if inp[i] == initial {
			continue
		}

		return false
	}
	fmt.Println(inp, "all equal")
	return true
}

func split_to_parts(inp string, ln int) []string {
	var res []string
	for i := 0; i <= len(inp) - ln; i += ln {
		res = append(res, inp[i:i + ln])
	}
	return res
}

func parse_range(first, second int) ([]int, error) {
	var res []int
	for i := first; i < second; i++ {
		stringified := strconv.Itoa(i)

		if len(stringified) % 2 != 0 {
			continue
		}

		/*max_slice := len(stringified) / 2
		for ; max_slice > 0; max_slice-- {
			if len(stringified) % max_slice != 0 {
				continue
			}
			split_parts := split_to_parts(stringified, max_slice)
			if are_equal(split_parts) {
				fmt.Println(max_slice, "---", split_parts, "---", stringified)
				res = append(res, i)
				break
			}
		}*/

		split_parts := split_to_parts(stringified, len(stringified)/2)
		if are_equal(split_parts) {
			res = append(res, i)
		}
	}
	return res, nil
}

func sum(inp []int) int {
	res := 0
	for i := range inp {
		res += inp[i]
	}
	return res
}

func get_answer(inp string) (int, error) {
	splitted := strings.Split(inp, ",")
	var invalid_ids []int
	for _, v := range splitted {
		s_val := strings.Split(v, "-")
		first, second := s_val[0], s_val[1]

		if len(first) % 2 != 0 && len(second) % 2 != 0 {
			continue
		}

		first_int, err := strconv.Atoi(first)
		if err != nil {
			return 0, err
		}

		second_int, err := strconv.Atoi(second)
		if err != nil {
			return 0, err
		}

		parse_res, err := parse_range(first_int, second_int)
		if err != nil {
			return 0, err
		}

		invalid_ids = append(invalid_ids, parse_res...)
	}

	return sum(invalid_ids), nil
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Couldn't read the input:", err)
		return
	}

	answer, err := get_answer(string(file))
	if err != nil {
		fmt.Println("Couldn't retrieve the answer:", err)
		return
	}

	fmt.Printf("Answer is: %v\n", answer)
}
