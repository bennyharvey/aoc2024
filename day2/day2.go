package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	u "github.com/bennyharvey/aoc2024/utils"
)

const (
	DIR_INC = iota
	DIR_DEC
)

type Report struct {
	numbers []int
	safe    bool
	dir     int
}

func (r *Report) setDir() {
	incCount := 0
	decCount := 0
	for i := 0; i < len(r.numbers)-1; i++ {
		if r.numbers[i] < r.numbers[i+1] {
			incCount++
		} else {
			decCount++
		}
	}
	if incCount > decCount {
		r.dir = DIR_INC
	} else {
		r.dir = DIR_DEC
	}
}

func SolvePart1() {
	file := u.Must(os.Open("day2/day2_test.txt"))
	defer file.Close()

	var reports []Report

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		numStrings := strings.Split(line, " ")
		var nums []int
		for _, num := range numStrings {
			nums = append(nums, u.Must(strconv.Atoi(num)))
		}
		reports = append(reports, Report{numbers: nums})
	}

	for j := 0; j < len(reports); j++ {
		r := &reports[j]
		r.setDir()
		for i := 0; i < len(r.numbers)-1; i++ {

			if r.dir == DIR_INC {
				if r.numbers[i] >= r.numbers[i+1] || r.numbers[i+1]-r.numbers[i] > 3 {
					r.safe = false
					break
				}
			}

			if r.dir == DIR_DEC {
				if r.numbers[i+1] >= r.numbers[i] || r.numbers[i]-r.numbers[i+1] > 3 {
					r.safe = false
					break
				}
			}
			r.safe = true
		}
	}

	count := 0
	for _, r := range reports {
		if r.safe {
			count++
		}
	}
	// u.PPrintSlice(reports)

	fmt.Println("Day 2, Part 1 Answer: ", count)
}

func SolvePart2(fileName string) int {

	// file := u.Must(os.Open("day2/day2_manual_test.txt"))
	file := u.Must(os.Open(fileName))
	// file := u.Must(os.Open("day2/day2_sample.txt"))
	defer file.Close()

	var reports []Report

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		numStrings := strings.Split(line, " ")
		var nums []int
		for _, num := range numStrings {
			nums = append(nums, u.Must(strconv.Atoi(num)))
		}
		reports = append(reports, Report{numbers: nums})
	}

	for j := 0; j < len(reports); j++ {
		r := &reports[j]
		r.setDir()
		fmt.Printf("%+v\n", r.numbers)
		// fmt.Printf("%+v\n", r.numbers[5])
		r.safe = sliceIsSafe(r.numbers, r.dir, false)
		fmt.Printf("Safe: %+v %v \n ", r.safe, r.dir)
		// fmt.Printf("%+v\n", r.numbers[5])
	}

	count := 0
	for _, r := range reports {
		if r.safe {
			count++
		}
	}

	// u.PPrintSlice(reports)

	fmt.Println("Day 2, Part 2 Answer: ", count)

	return count

}

func sliceIsSafe(s []int, dir int, hadError bool) bool {
	log.Printf("Checking slice %v", s)
	for i := 0; i < len(s)-1; i++ {
		if dir == DIR_INC {
			if s[i] >= s[i+1] || s[i+1]-s[i] > 3 {
				if i == 0 && !hadError {
					t1 := make([]int, len(s))
					t2 := make([]int, len(s))
					copy(t1, s)
					copy(t2, s)
					return sliceIsSafe(remove(t1, i), dir, true) || sliceIsSafe(remove(t2, i+1), dir, true)
				}
				if !hadError {
					return sliceIsSafe(remove(s, i+1), dir, true)
				}
				return false
			}
		}

		if dir == DIR_DEC {
			if s[i+1] >= s[i] || s[i]-s[i+1] > 3 {
				if i == 0 && !hadError {
					t1 := make([]int, len(s))
					t2 := make([]int, len(s))
					copy(t1, s)
					copy(t2, s)
					return sliceIsSafe(remove(t1, i), dir, true) || sliceIsSafe(remove(t2, i+1), dir, true)
				}
				if !hadError {
					return sliceIsSafe(remove(s, i+1), dir, true)
				}
				return false
			}
		}
	}
	return true
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
