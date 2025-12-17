package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

var operations = map[string](func(int, int) int){
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
}

func count_slice(inp []int, op string) int {
	res := inp[0]
	for i := 1; i < len(inp); i++ {
		res = operations[op](res, inp[i])
	}
	return res
}

func rebuild_slice(inp [][]string) [][]string {
	res := make([][]string, len(inp[0]))
	for i := range inp {
		for j := range inp[i] {
			res[j] = append(res[j], inp[i][j])
		}
	}
	return res
}

func convert_slice(inp [][]string) [][]int {
	new_int_slice := make([][]int, len(inp))
	for i := range inp {
		for j := range inp[i] {
			new_val, _ := strconv.Atoi(inp[i][j])
			new_int_slice[i] = append(new_int_slice[i], new_val)
		}
	}
	return new_int_slice
}

func conv_and_count(inp string) int {
	elem_line_len := len(strings.Fields(strings.Split(inp, "\n")[0]))
	fields := strings.Fields(inp)

	var new_s_str [][]string
	for i := elem_line_len; i <= len(fields); i += elem_line_len {
		new_s_str = append(new_s_str, fields[i-elem_line_len:i])
	}

	ops := new_s_str[len(new_s_str)-1]

	new_s_str = new_s_str[:len(new_s_str)-1]
	new_s_str = rebuild_slice(new_s_str)

	new_int_slice := convert_slice(new_s_str)

	var subtotal int
	for j := range new_int_slice {
		subtotal += count_slice(new_int_slice[j], ops[j])
	}

	return subtotal
}

func get_answer(inp string) int {
	return conv_and_count(inp)
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Couldn't open file:", err)
		return
	}

	answer := get_answer(string(file))

	fmt.Println("The answer is:", answer)
}
