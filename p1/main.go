package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed p1.1data.txt
var p1Data embed.FS

func part1() {
	data, err := p1Data.ReadFile("p1.1data.txt")
	if err != nil {
		panic(err)
	}
	entries := strings.Split(string(data), "\n")
	col1 := []int{}
	col2 := []int{}

	for _, entry := range entries {
		parts := strings.Split(entry, "   ")

		c1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		c2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		col1 = append(col1, c1)
		col2 = append(col2, c2)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	diffs := []int{}
	for i, c1 := range col1 {
		c2 := col2[i]
		diffs = append(diffs, abs(c1-c2))
	}

	sum := 0
	for _, val := range diffs {
		sum += val
	}

	fmt.Println(sum)
}

func part2() {
	data, err := p1Data.ReadFile("p1.1data.txt")
	if err != nil {
		panic(err)
	}

	his1 := map[int]int{}
	his2 := map[int]int{}

	entries := strings.Split(string(data), "\n")
	for _, entry := range entries {
		parts := strings.Split(entry, "   ")

		c1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		his1[c1] += 1

		c2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		his2[c2] += 1
	}

	sum := 0
	for k, val := range his1 {
		val2 := his2[k]

		sum += k * val * val2
	}

	fmt.Println(sum)
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func main() {
	part2()
}
