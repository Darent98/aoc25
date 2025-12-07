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

	solve(exampleRows)
	solve(inputRows)

	solve2(exampleRows)
	solve2(inputRows)

}

func solve(rows []string) {
	res := 0

	fmt.Println("The result is:", res)
}

func solve2(rows []string) {
	res := 0

	fmt.Println("The result is:", res)
}
