package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test1.txt")
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

	r := regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)
	var result1 int = 0
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			l, _ := strconv.Atoi(v[1])
			r, _ := strconv.Atoi(v[2])
			result1 += l * r
		}
	}
	fmt.Println("1:", result1)
}

func part2() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("test2.txt")
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

	r := regexp.MustCompile(`mul\((\d+)\,(\d+)\)|do\(\)|don't\(\)`)
	var result int = 0
	canDo := true
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			if v[0][:3] == "mul" && canDo {
				l, _ := strconv.Atoi(v[1])
				r, _ := strconv.Atoi(v[2])
				result += l * r
			} else if v[0] == "do()" {
				canDo = true
			} else if v[0] == "don't()" {
				canDo = false
			}
		}
	}
	fmt.Println("2:", result)

}
func main() {
	part1()
	part2()
}
