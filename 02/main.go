package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(fields []string) bool {
	l, _ := strconv.Atoi(fields[0])
	r, _ := strconv.Atoi(fields[1])
	isInc := r > l
	for i := 0; i < len(fields)-1; i++ {
		l, _ := strconv.Atoi(fields[i])
		r, _ := strconv.Atoi(fields[i+1])
		var diff int
		if isInc {
			diff = r - l
		} else {
			diff = l - r
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func part2(fields []string) bool {
	if part1(fields) {
		return true
	}

	for i := range len(fields) {
		var fields2 []string
		for j, f := range fields {
			if i == j {
				continue
			}
			fields2 = append(fields2, f)
		}
		if part1(fields2) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var lines_safe1 int = 0
	for _, line := range lines {
		fields := strings.Fields(line)
		safe := part1(fields)
		if safe {
			lines_safe1 += 1
		}
	}
	fmt.Println("1:", lines_safe1)

	var lines_safe2 int = 0
	for _, line := range lines {
		fields := strings.Fields(line)
		safe := part2(fields)
		if safe {
			lines_safe2 += 1
		}
	}
	fmt.Println("2:", lines_safe2)

}
