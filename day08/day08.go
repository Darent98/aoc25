package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Point struct {
	p1 int
	p2 int
	p3 int
}

type Distance struct {
	distance float64
	index1   int
	index2   int
}

func main() {
	exampleRows := strings.Split(example, "\n")
	inputRows := strings.Split(input, "\n")

	solve(exampleRows, 10)
	solve(inputRows, 1000)

	solve2(exampleRows)
	solve2(inputRows)

}

func solve(rows []string, size int) {
	res := 0

	junctions := []Point{}

	closestDistances := make([]Distance, 0, size)

	circuits := [][]int{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		junctions = append(junctions, getPoint(row))
	}

	for i := range len(junctions) {
		for j := i + 1; j < len(junctions); j++ {
			d := calcDistance(junctions[i], junctions[j])
			if len(closestDistances) < size {
				closestDistances = append(closestDistances, Distance{distance: d, index1: i, index2: j})
			} else {
				sort.Slice(closestDistances, func(a, b int) bool {
					return closestDistances[a].distance < closestDistances[b].distance
				})
				if d < closestDistances[len(closestDistances)-1].distance {
					closestDistances = append(closestDistances[:len(closestDistances)-1], Distance{distance: d, index1: i, index2: j})
				}
			}
		}
	}

	for _, cD := range closestDistances {
		indexIncluded1 := findIndex(circuits, cD.index1)
		indexIncluded2 := findIndex(circuits, cD.index2)

		if indexIncluded1 == nil && indexIncluded2 == nil {
			circuits = append(circuits, []int{cD.index1, cD.index2})
			continue
		}
		if indexIncluded1 != nil && indexIncluded2 == nil {
			circuits[*indexIncluded1] = append(circuits[*indexIncluded1], cD.index2)
			continue
		}
		if indexIncluded1 == nil && indexIncluded2 != nil {
			circuits[*indexIncluded2] = append(circuits[*indexIncluded2], cD.index1)
			continue
		}
		if *indexIncluded1 == *indexIncluded2 {
			continue
		}
		if indexIncluded1 != nil && indexIncluded2 != nil {
			circuits[*indexIncluded1] = append(circuits[*indexIncluded1], circuits[*indexIncluded2]...)
			circuits = append(circuits[:*indexIncluded2], circuits[*indexIncluded2+1:]...)
			continue
		}
	}

	sort.Slice(circuits, func(a, b int) bool {
		return len(circuits[a]) > len(circuits[b])
	})

	res = len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	fmt.Println("The result is:", res)
}

func solve2(rows []string) {
	res := 0

	junctions := []Point{}

	closestDistances := make([]Distance, 0)

	circuits := [][]int{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		junctions = append(junctions, getPoint(row))
	}

	for i := range len(junctions) {
		for j := i + 1; j < len(junctions); j++ {
			d := calcDistance(junctions[i], junctions[j])
			closestDistances = append(closestDistances, Distance{distance: d, index1: i, index2: j})
		}
	}

	sort.Slice(closestDistances, func(a, b int) bool {
		return closestDistances[a].distance < closestDistances[b].distance
	})

	for _, cD := range closestDistances {
		indexIncluded1 := findIndex(circuits, cD.index1)
		indexIncluded2 := findIndex(circuits, cD.index2)
		matched := false

		if indexIncluded1 == nil && indexIncluded2 == nil {
			circuits = append(circuits, []int{cD.index1, cD.index2})
			matched = true
		}
		if !matched && indexIncluded1 != nil && indexIncluded2 == nil {
			circuits[*indexIncluded1] = append(circuits[*indexIncluded1], cD.index2)
			matched = true
		}
		if !matched && indexIncluded1 == nil && indexIncluded2 != nil {
			circuits[*indexIncluded2] = append(circuits[*indexIncluded2], cD.index1)
			matched = true
		}
		if !matched && *indexIncluded1 == *indexIncluded2 {
			matched = true
		}
		if !matched && indexIncluded1 != nil && indexIncluded2 != nil {
			circuits[*indexIncluded1] = append(circuits[*indexIncluded1], circuits[*indexIncluded2]...)
			circuits = append(circuits[:*indexIncluded2], circuits[*indexIncluded2+1:]...)
			matched = true
		}

		if len(circuits) == 1 && len(circuits[0]) == len(junctions) {
			res = junctions[cD.index1].p1 * junctions[cD.index2].p1
			break
		}

	}

	fmt.Println("The result is:", res)
}

func calcDistance(p Point, q Point) float64 {
	var val float64 = math.Pow(float64(p.p1-q.p1), 2) + math.Pow(float64(p.p2-q.p2), 2) + math.Pow(float64(p.p3-q.p3), 2)
	return math.Sqrt(val)
}

func getPoint(row string) Point {
	splitted := strings.Split(row, ",")
	p1, err := strconv.Atoi(splitted[0])
	if err != nil {
		panic(err)
	}
	p2, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}
	p3, err := strconv.Atoi(splitted[2])
	if err != nil {
		panic(err)
	}
	return Point{
		p1: p1,
		p2: p2,
		p3: p3,
	}
}

func findIndex(data [][]int, target int) *int {
	for i, slice := range data {
		for _, v := range slice {
			if v == target {
				return &i
			}
		}
	}
	return nil
}
