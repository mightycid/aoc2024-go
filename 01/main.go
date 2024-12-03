package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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
		// fmt.Println(line)
		lines = append(lines, line)
	}
	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	var l, r []int
	for _, line := range lines {
		splits := strings.Split(line, "   ")
		ln, _ := strconv.Atoi(splits[0])
		rn, _ := strconv.Atoi(splits[1])
		l = append(l, ln)
		r = append(r, rn)
	}

	l1 := make([]int, len(l))
	r1 := make([]int, len(r))
	copy(l1, l)
	copy(r1, r)

	slices.Sort(l1)
	slices.Sort(r1)

	var sumDiff int = 0
	for i := 0; i < len(l1); i++ {
		sumDiff += int(math.Abs(float64(l1[i]) - float64(r1[i])))
	}

	fmt.Println("1:", sumDiff)

	r2 := make([]int, len(r))
	copy(r2, r)

	counter := map[int]int{}
	for _, e := range r2 {
		counter[e] += 1
	}
	var simScore int = 0
	for _, e := range l {
		c, ok := counter[e]
		if ok {
			simScore += e * c
		}
	}
	fmt.Println("2:", simScore)
}
