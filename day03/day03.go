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

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		indexHighest := indexOfHighestNumber([]rune(row[:len(row)-1]))
		indexSecondHighest := indexOfHighestNumber([]rune(row[indexHighest+1:])) + indexHighest + 1

		s := string(row[indexHighest]) + string(row[indexSecondHighest])
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res += n
	}

	fmt.Println("The result is:", res)
}

func solve2(rows []string) {
	res := 0

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		s := ""

		startIndex := 0
		for i := 11; i >= 0; i-- {
			startIndex = indexOfHighestNumber([]rune(row[startIndex:len(row)-i])) + 1 + startIndex
			s += string(row[startIndex-1])
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res += n
	}

	fmt.Println("The result is:", res)
}

func indexOfHighestNumber(numbers []rune) int {
	indexHighest := 0
	highest := numbers[indexHighest]
	for i := 0; i < (len(numbers)); i++ {
		if numbers[i] > highest {
			highest = numbers[i]
			indexHighest = i
		}
	}
	return indexHighest
}

func runeToInt(r rune) int {
	return int(r - '0')
}
