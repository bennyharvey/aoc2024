package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/bennyharvey/aoc2024/utils"
)

func SolvePart1() {
	file := utils.Must(os.Open("day1/day1_test.txt"))
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "   ")
		left = append(left, utils.Must(strconv.Atoi(s[0])))
		right = append(right, utils.Must(strconv.Atoi(s[1])))
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += utils.Abs(left[i] - right[i])
	}

	fmt.Println("Day 1, Part 1 Answer: ", sum)
}

func SolvePart2() {
	file := utils.Must(os.Open("day1/day1_test.txt"))
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "   ")
		left = append(left, utils.Must(strconv.Atoi(s[0])))
		right = append(right, utils.Must(strconv.Atoi(s[1])))
		// fmt.Println(line)
	}

	rightMap := make(map[int]int)
	for _, n := range right {
		_, ok := rightMap[n]
		if !ok {
			rightMap[n] = 0
		}
		rightMap[n] += 1
	}

	sum := 0
	for _, n := range left {
		count, ok := rightMap[n]
		if !ok {
			continue
		}
		sum += count * n
	}

	fmt.Println("Day 1, Part 2 Answer: ", sum)
}
