package day5

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	u "github.com/bennyharvey/aoc2024/utils"
)

type Manual struct {
	pages   []int
	correct bool
}



func SolvePart1(fileName string) {
	manuals, _ := parseManuals(fileName)
	sum := 0
	for _, manual := range manuals {
		if manual.correct {
			sum += manual.pages[len(manual.pages)/2]
		}
		u.Printf1("%+v\n", manual)
	}
	u.Println("Day 5, Part 1: ", sum)
}

func SolvePart2(fileName string) {
	manuals, rules := parseManuals(fileName)
	sum := 0
	for mi := 0; mi < len(manuals); mi++ {
		if !manuals[mi].correct {
			sort.Slice(manuals[mi].pages, func(i, j int) bool {
				return slices.Contains(rules[manuals[mi].pages[j]], manuals[mi].pages[i])
			})
			sum += manuals[mi].pages[len(manuals[mi].pages)/2]
		}
	}
	u.Println("Day 5, Part 2: ", sum)

}

func parseManuals(fileName string) ([]Manual, map[int][]int) {
	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)

	high := make(map[int][]int)
	low := make(map[int][]int)
	manuals := make([]Manual, 0)

	stage := 1
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		u.Println1(line)
		if line == "" {
			stage = 2
			continue
		}
		if stage == 1 {
			parts := strings.Split(line, "|")
			left := u.Must(strconv.Atoi(parts[0]))
			right := u.Must(strconv.Atoi(parts[1]))
			high[left] = append(high[left], right)
			low[right] = append(low[right], left)
		}
		if stage == 2 {
			pageNumbers := strings.Split(line, ",")
			manuals = append(manuals, Manual{correct: true})
			for _, pageNumber := range pageNumbers {
				manuals[i].pages = append(manuals[i].pages, u.Must(strconv.Atoi(pageNumber)))
			}
			i++
		}
	}

	for mi := 0; mi < len(manuals); mi++ {
		manual := &manuals[mi]
		for i := 0; i < len(manual.pages); i++ {
			for leftPart := 0; leftPart < i; leftPart++ {
				if slices.Contains(high[manual.pages[i]], manual.pages[leftPart]) {
					manual.correct = false
				}
			}
		}
	}

	return manuals, low
}
