package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var dataFile embed.FS

const (
	unknown int = iota
	inc
	dec
)

func part1() {
	data, err := dataFile.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(data), "\n")
	sum := 0
	for _, entry := range entries {
		level := 0
		parts := strings.Split(entry, " ")

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		safe := true
		for _, b := range parts[1:] {
			num2, err := strconv.Atoi(b)
			if err != nil {
				panic(err)
			}

			diff := num2 - num1
			if diff == 0 {
				safe = false
				break
			}

			if diff > 3 || diff < -3 {
				safe = false
				break
			}

			if level*diff < 0 {
				safe = false
				break
			}

			level = diff
			num1 = num2
		}

		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}

func isSafe(level, num1 int, skip int, parts []string) (bool, int) {
	for i := 1; i < len(parts); i++ {
		if i == skip {
			continue
		}

		b := parts[i]
		num2, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		diff := num2 - num1
		if diff == 0 {
			return false, i
		}

		if diff > 3 || diff < -3 {
			return false, i
		}

		if level*diff < 0 {
			return false, i
		}

		level = diff
		num1 = num2
	}

	return true, 0
}

func part2() {
	data, err := dataFile.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(data), "\n")
	sum := 0
	for _, entry := range entries {
		level := 0
		parts := strings.Split(entry, " ")

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		safe, skip := isSafe(level, num1, -1, parts)
		if !safe {
			fmt.Println(parts, skip)
			safe, _ = isSafe(level, num1, skip-1, parts)
			fmt.Println("is it safe after skipping", parts, safe)
		}

		if safe {
			sum++
		}
	}

	fmt.Println(sum)
}

func main() {
	part2()
}
