package day3

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	u "github.com/bennyharvey/aoc2024/utils"
)

func SolvePart1() {
	file := u.Must(os.Open("day3/day3_test.txt"))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text string
	sum := 0
	dexp := regexp.MustCompile(`\d{1,3}`)
	exp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for scanner.Scan() {
		text = scanner.Text()
		u.Println3(text)
		matches := exp.FindAllString(text, -1)
		for _, match := range matches {
			u.Println3(match)
			numbers := dexp.FindAllString(match, -1)
			mul := 1
			for _, number := range numbers {
				mul *= u.Must(strconv.Atoi(number))
			}
			u.Println2(mul)
			sum += mul
		}
	}

	u.Println("Day 3, Part 1 Answer: ", sum)

}

func SolvePart2() {
	file := u.Must(os.Open("day3/day3_test.txt"))
	defer file.Close()

	sum := 0
	dexp := regexp.MustCompile(`\d{1,3}`)
	exp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	state := "do"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		matches := exp.FindAllString(text, -1)
		for _, match := range matches {
			u.Println2(match)
			if match == "do()" {
				state = "do"
				continue
			}
			if match == "don't()" {
				state = "dont"
				continue
			}
			if state == "dont" {
				continue
			}
			numbers := dexp.FindAllString(match, -1)
			mul := 1
			for _, number := range numbers {
				mul *= u.Must(strconv.Atoi(number))
			}
			sum += mul
			u.Println2(mul, sum)
		}
	}
	u.Println("Day 3, Part 2 Answer: ", sum)

}
