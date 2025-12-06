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

	example = strings.ReplaceAll(example, "\n", "")
	input = strings.ReplaceAll(input, "\n", "")

	solve(strings.Split(example, ","))
	solve(strings.Split(input, ","))

	//solve2([]string{"824824821-824824827"})
	solve2(strings.Split(example, ","))
	solve2(strings.Split(input, ","))
}

func solve(ranges []string) {
	res := 0

	for _, r := range ranges {
		tmp := strings.Split(r, "-")
		start, err := strconv.Atoi(tmp[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}

		for start <= end {
			n := strconv.Itoa(start)
			runes := []rune(n)
			length := len(runes)
			if (length % 2) != 0 {
				start++
				continue
			}

			mid := length / 2
			//left := runes[:mid]
			//right := runes[mid:]

			equal := true
			for i := range mid {
				if runes[i] != runes[mid+i] {
					equal = false
					break
				}
			}
			if equal {
				res += start
			}

			start++
		}

	}

	fmt.Println("The result is:", strconv.Itoa(res))
}

func solve2(ranges []string) {
	res := 0

	for _, r := range ranges {
		tmp := strings.Split(r, "-")
		start, err := strconv.Atoi(tmp[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}

		for start <= end {
			n := strconv.Itoa(start)
			runes := []rune(n)
			length := len(runes)

			// equal is here false
			equal := false

			//fmt.Println("number, length:", n, length)
			// create splits
			for splitLength := 1; splitLength <= (length / 2); splitLength++ {
				//fmt.Println("split length:", splitLength)
				changeEqual := true
				if length%splitLength != 0 {
					continue
				}

				splits := length / splitLength
				//fmt.Println("splits:", splits)
				// iterate splits
				for i := 1; i < splits; i++ {

					stop := false
					// compare first split with selected split
					for k := 0; k < splitLength; k++ {
						if runes[k] != runes[i*splitLength+k] {
							changeEqual = false
							stop = true
							break
						}
					}
					if stop {
						break
					}
				}
				if changeEqual {
					equal = true
					break
				}
			}
			if equal {
				res += start
			}

			start++
		}
	}

	fmt.Println("The result is:", strconv.Itoa(res))
}
