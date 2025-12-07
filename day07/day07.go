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

	exampleRows = strings.Split(example, "\n")
	inputRows = strings.Split(input, "\n")
	solve2(exampleRows)
	solve2(inputRows)

}

func solve(rows []string) {
	res := 0

	for i := range rows {
		if i == 0 {
			continue
		}
		for k := range rows[i] {
			runes := []rune(rows[i])
			switch cell := rune(rows[i-1][k]); cell {
			case 'S':
				runes[k] = '|'
			case '|':
				if runes[k] == '.' {
					runes[k] = '|'
				}
				if runes[k] == '^' {
					if k != 0 && runes[k-1] == '.' {
						runes[k-1] = '|'
					}
					if k < len(runes) && runes[k+1] == '.' {
						runes[k+1] = '|'
					}
				}
			default:
			}
			rows[i] = string(runes)

			if rows[i][k] == '^' && rows[i-1][k] == '|' {
				res++
			}
		}
		// fmt.Println(rows[i])
	}

	fmt.Println("The result is:", res)
}

func solve2(rows []string) {
	res := 0

	arr := [][]string{}
	for i := range rows[:len(rows)-1] {
		arr = append(arr, stringToStringSlice(rows[i]))
	}
	for i := range arr {
		if i == 0 {
			continue
		}
		//fmt.Println(arr[i])
		for k := range arr[i] {
			switch cell := arr[i-1][k]; cell {
			case "S":
				arr[i][k] = "1"
			case "^":
				continue
			case ".":
				continue
			default:
				// i am dot, above me number
				if arr[i][k] == "." {
					arr[i][k] = cell
					continue
				}
				// i am splitter, above me number
				if arr[i][k] == "^" {
					n, err := strconv.Atoi(cell)
					if err != nil {
						panic(err)
					}
					if k != 0 {
						if arr[i][k-1] == "^" {
							continue
						}
						if arr[i][k-1] == "." {
							arr[i][k-1] = cell
						} else {
							nLeft, err := strconv.Atoi(arr[i][k-1])
							if err != nil {
								panic(err)
							}
							nLeft += n
							arr[i][k-1] = strconv.Itoa(nLeft)
						}
					}
					if k < len(arr[i]) {
						if arr[i][k+1] == "^" {
							continue
						}
						if arr[i][k+1] == "." {
							arr[i][k+1] = cell
						} else {
							nRight, err := strconv.Atoi(arr[i][k+1])
							if err != nil {
								panic(err)
							}
							nRight += n
							arr[i][k+1] = strconv.Itoa(nRight)
						}
					}
					continue
				}
				// i am number, above me is also number
				if arr[i][k] != "." && arr[i][k] != "^" {
					nc, err := strconv.Atoi(cell)
					if err != nil {
						panic(err)
					}
					n, err := strconv.Atoi(arr[i][k])
					if err != nil {
						panic(err)
					}
					arr[i][k] = strconv.Itoa(n + nc)
				}
			}
		}
		//fmt.Println(arr[i])
	}

	for _, cell := range arr[len(arr)-1] {
		if cell != "." {
			n, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			res += n
		}
	}

	fmt.Println("The result is:", res)
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func stringToStringSlice(s string) []string {
	runes := []rune(s)
	out := make([]string, len(runes))
	for i, r := range runes {
		out[i] = string(r)
	}
	return out
}
