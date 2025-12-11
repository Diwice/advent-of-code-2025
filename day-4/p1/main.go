package main

import (
	"os"
	"fmt"
	"strings"
)

func check_adjucents(mx [][]bool, w, h, c_w, c_h int) bool {
	count := 0

	neighbors := make([]bool, 8)

	if c_w > 0 {
		neighbors[3] = mx[c_h][c_w - 1]

		if c_h > 0 {
			neighbors[0] = mx[c_h - 1][c_w - 1]
		}

		if c_h < h - 1 {
			neighbors[5] = mx[c_h + 1][c_w - 1]
		}
	}

	if c_w < w - 1 {
		neighbors[4] = mx[c_h][c_w + 1]

		if c_h > 0 {
			neighbors[2] = mx[c_h - 1][c_w + 1]
		}

		if c_h < h - 1 {
			neighbors[7] = mx[c_h + 1][c_w + 1]
		}
	}

	if c_h > 0 {
		neighbors[1] = mx[c_h - 1][c_w]
	}

	if c_h < h - 1 {
		neighbors[6] = mx[c_h + 1][c_w]
	}

	for i := range neighbors {
		if neighbors[i] {
			count++
		}
	}

	return count < 4
}

func count_rolls(mx [][]bool, w, h int) int {
	res := 0
	for cur_h := range mx {
		for cur_w := range mx[cur_h] {
			if mx[cur_h][cur_w] && check_adjucents(mx, w, h, cur_w, cur_h) {
				res++
			}
		}
	}
	return res
}

func fill_matrix(mx *[][]bool, line string, elem int) {
	for i, v := range line {
		if v == '@' {
			(*mx)[elem][i] = true
		}
	}
}

func make_matrix(w, h int) [][]bool {
	new_matrix := make([][]bool, h)
	for i := range new_matrix {
		new_matrix[i] = make([]bool, w)
	}
	return new_matrix
}

func get_answer(inp string) int {
	splitted := strings.Split(inp, "\n")

	width, height := len(strings.TrimSpace(splitted[0])), len(splitted)
	matrix := make_matrix(width, height)

	for i, v := range splitted {
		splitted[i] = strings.TrimSpace(v)

		fill_matrix(&matrix, splitted[i], i)
	}

	res := count_rolls(matrix, width, height)

	return res
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Couldn't read file:", err)
		return
	}

	answer := get_answer(string(file))

	fmt.Println("The answer is:", answer)
}
