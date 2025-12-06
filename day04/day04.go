package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	exampleRows := strings.Split(example, "\n")
	inputRows := strings.Split(input, "\n")

	solve(exampleRows, true)
	solve(inputRows, true)

	solve2(exampleRows)
	solve2(inputRows)

}

func solve(rows []string, printRes bool) int {
	res := 0

	for i, row := range rows {
		r := []rune(rows[i])
		for ir, cell := range row {
			if cell == '.' {
				continue
			}
			if canAccess(i, ir, rows) {
				res++
				r[ir] = 't'
			}
		}
		rows[i] = string(r)
	}

	if printRes {
		fmt.Println("The result is:", res)
	}
	return res
}

func solve2(rows []string) {
	res := 0

	for true {
		n := solve(rows, false)
		if n == 0 {
			break
		}
		res += n
		for i := range rows {
			rows[i] = strings.ReplaceAll(rows[i], "t", ".")
		}
	}

	fmt.Println("The result is:", res)
}

func canAccess(i int, ir int, rows []string) bool {
	adjacent := 0
	for k := i - 1; k <= i+1; k++ {
		if k < 0 || k > len(rows)-1 {
			continue
		}
		for kr := ir - 1; kr <= ir+1; kr++ {
			if adjacent > 3 {
				return false
			}
			if kr < 0 || kr > len(rows[k])-1 {
				continue
			}
			// skip if this is the center
			if k == i && kr == ir {
				continue
			}
			if rows[k][kr] == '@' || rows[k][kr] == 't' {
				adjacent++
			}
		}
	}
	if adjacent > 3 {
		return false
	}
	return true
}
