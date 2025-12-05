package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	res := 0

	// open file
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := `([LR]{1})(\d{1,3})$`

	start := 50

	for scanner.Scan() {
		line := scanner.Text()

		matches := regexp.MustCompile(reg).FindStringSubmatch(line)
		direction := matches[1]
		clicks, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		// fmt.Println(matches[1] + " " + matches[2])

		if direction == "L" {
			clicks = clicks * -1
		}

		startBefore := start
		start += clicks

		// start of extension for part two
		timesCrossed0 := 0
		if (start < 0 && startBefore > 0) || (start > 0 && startBefore < 0) {
			timesCrossed0++
		}
		timesCrossed0 += (abs(start) / 100)
		if start != 0 && start%100 == 0 {
			timesCrossed0--
		}
		// fmt.Println(startBefore, start, timesCrossed0)
		res += timesCrossed0

		// end

		start = start % 100

		if start == 0 {
			res++
		}
	}

	fmt.Println("The result is: " + strconv.Itoa(res))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
