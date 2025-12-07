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

	getRanges := true
	ranges := []string{}

	for _, row := range rows {
		if len(row) < 1 {
			getRanges = false
			continue
		}
		if getRanges {
			ranges = append(ranges, row)
			/*
				tmp := strings.Split(row, "-")
				start, err := strconv.Atoi(tmp[0])
				if err != nil {
					panic(err)
				}
				end, err := strconv.Atoi(tmp[1])
				if err != nil {
					panic(err)
				}
					for i := start; i <= end; i++ {
						fresh[i] = true
					}
			*/

		} else {
			id, err := strconv.Atoi(row)
			if err != nil {
				panic(err)
			}
			/*
				_, ok := fresh[id]
				if ok {
					res++
				}
			*/
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
				if start <= id && id <= end {
					res++
					break
				}
			}
		}
	}

	fmt.Println("The result is:", res)
}

type Range struct {
	start int
	end   int
}

func solve2(rows []string) {
	res := 0
	fresh := []Range{}
	for _, row := range rows {
		if len(row) < 1 {
			break
		}
		tmp := strings.Split(row, "-")
		start, err := strconv.Atoi(tmp[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}

		if len(fresh) == 0 {
			fresh = append(fresh, Range{start: start, end: end})
			continue
		}

		notMatched := true
		for r := range fresh {
			notMatched = !matchAndUpdate(fresh, r, start, end)
			if !notMatched {
				break
			}
		}
		if notMatched {
			fresh = append(fresh, Range{start: start, end: end})
		}
	}

	for true {
		match := false
		for i, r := range fresh {
			if i == len(fresh)-1 {
				break
			}
			for k := i + 1; k < len(fresh); k++ {
				match = matchAndUpdate(fresh, k, r.start, r.end)
				if match {
					fresh = append(fresh[:i], fresh[i+1:]...)
					break
				}
			}
			if match {
				break
			}
		}
		if match {
			continue
		}
		break
	}

	for _, r := range fresh {
		res += (r.end - r.start) + 1
	}

	fmt.Println("The result is:", res)
}

func matchAndUpdate(fresh []Range, r int, start int, end int) bool {
	// range is in existing range
	if fresh[r].start <= start && fresh[r].end >= end {
		return true
	}
	// range includes existing range
	if fresh[r].start >= start && fresh[r].end <= end {
		fresh[r].start = start
		fresh[r].end = end
		return true
	}
	// range overlaps at start with existing range
	if fresh[r].start > start && fresh[r].end >= end && fresh[r].start <= end {
		fresh[r].start = start
		return true
	}
	// range overlaps at end with existing range
	if fresh[r].start <= start && fresh[r].end < end && start <= fresh[r].end {
		fresh[r].end = end
		return true
	}
	return false
}
