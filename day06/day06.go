package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	exampleRows := strings.Split(example, "\n")
	inputRows := strings.Split(input, "\n")

	solve(exampleRows)
	solve(inputRows)

	solve2(exampleRows)
	solve2(inputRows)

}

func solve(rows []string) {
	res := 0

	matrix := [][]int{}
	operators := []rune{}

	for i, row := range rows {
		for cell := range strings.SplitSeq(row, " ") {
			if len(cell) == 0 {
				continue
			}
			if (len(rows) - 2) == i {
				operators = append(operators, rune(cell[0]))
				continue
			}
			n, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			if len(matrix) == i {
				matrix = append(matrix, []int{n})
				continue
			}
			matrix[i] = append(matrix[i], n)
		}
	}

	res = solveMatrix(operators, matrix)

	fmt.Println("The result is:", res)
}

func solve2(rows []string) {
	res := 0

	matrix := [][]int{}
	tmpMatrix := []string{}
	operators := []rune{}

	for i, row := range rows {
		for k, r := range row {
			if (len(rows) - 2) == i {
				if r != ' ' {
					operators = append(operators, r)
				}
				continue
			}
			if len(tmpMatrix) == k {
				tmpMatrix = append(tmpMatrix, string(r))
				continue
			}
			tmpMatrix[k] = fmt.Sprintf("%s%s", tmpMatrix[k], string(r))
		}
	}

	i := 0
	for _, row := range tmpMatrix {
		replaced := strings.ReplaceAll(row, " ", "")
		if len(replaced) == 0 {
			i++
			continue
		}
		n, err := strconv.Atoi(replaced)
		if err != nil {
			panic(err)
		}
		if len(matrix) == i {
			matrix = append(matrix, []int{n})
			continue
		}
		matrix[i] = append(matrix[i], n)
	}

	res = solveMatrix2(operators, matrix)

	fmt.Println("The result is:", res)
}

func solveMatrix(operators []rune, matrix [][]int) int {
	res := 0
	for i := range len(operators) {
		tmp := 0
		if operators[i] == '*' {
			tmp = 1
		}
		for k := range len(matrix) {
			if operators[i] == '*' {
				tmp *= matrix[k][i]
			} else {
				tmp += matrix[k][i]
			}
		}
		res += tmp
	}
	return res
}

func solveMatrix2(operators []rune, matrix [][]int) int {
	res := 0
	for i := range len(operators) {
		tmp := 0
		if operators[i] == '*' {
			tmp = 1
		}

		for k := range len(matrix[i]) {
			if operators[i] == '*' {
				tmp *= matrix[i][k]
			} else {
				tmp += matrix[i][k]
			}
		}
		res += tmp
	}
	return res
}
