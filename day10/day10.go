package day10

import (
	"bufio"
	"os"

	u "github.com/bennyharvey/aoc2024/utils"
)

type TrailStart struct {
	row, col int
}

type Field struct {
	rows [][]rune
}

func (f *Field) locateInt(row, col int) int {
	if row < 0 || col < 0 || row >= len(f.rows) || col >= len(f.rows[row]) {
		return -99
	}
	return int(f.rows[row][col] - '0')
}

type Location struct {
	row, col int
}

type Finishes struct {
	items map[int]Location
}

func (f *Field) findDistinctFinished(row, col int) (Finishes, bool) {
	UP := []int{-1, 0}
	DOWN := []int{1, 0}
	LEFT := []int{0, -1}
	RIGHT := []int{0, 1}
	DIRECTIONS := [][]int{UP, DOWN, LEFT, RIGHT}

	finishes := Finishes{items: make(map[int]Location)}

	currentStep := f.locateInt(row, col)
	for _, dir := range DIRECTIONS {
		drow := dir[0]
		dcol := dir[1]
		nextStep := f.locateInt(row+drow, col+dcol)
		if nextStep-currentStep == 1 && nextStep == 9 {
			coord := (row+drow)*10 + col + dcol
			finishes.items[coord] = Location{row + drow, col + dcol}
			continue
		}
		if nextStep-currentStep == 1 {
			nextFinishes, found := f.findDistinctFinished(row+drow, col+dcol)
			if found {
				for _, item := range nextFinishes.items {
					coord := item.row*10 + item.col
					finishes.items[coord] = Location{item.row, item.col}
				}
			}
		}
	}

	if len(finishes.items) > 0 {
		return finishes, true
	} else {
		return finishes, false
	}

}

func (f *Field) findAllFinishes(row, col int) ([]Location, bool) {
	UP := []int{-1, 0}
	DOWN := []int{1, 0}
	LEFT := []int{0, -1}
	RIGHT := []int{0, 1}
	DIRECTIONS := [][]int{UP, DOWN, LEFT, RIGHT}

	finishes := make([]Location, 0)

	currentStep := f.locateInt(row, col)
	for _, dir := range DIRECTIONS {
		drow := dir[0]
		dcol := dir[1]
		nextStep := f.locateInt(row+drow, col+dcol)
		if nextStep-currentStep == 1 && nextStep == 9 {

			finishes = append(finishes, Location{row + drow, col + dcol})
			continue
		}
		if nextStep-currentStep == 1 {
			nextFinishes, found := f.findAllFinishes(row+drow, col+dcol)
			if found {
				for _, item := range nextFinishes {
					finishes = append(finishes, Location{item.row, item.col})
				}
			}
		}
	}

	if len(finishes) > 0 {
		return finishes, true
	} else {
		return finishes, false
	}

}

func SolvePart1(fileName string) int {

	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)

	field := Field{rows: make([][]rune, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		field.rows = append(field.rows, []rune(line))
		u.Println3(line)
	}

	starts := make([]TrailStart, 0)
	for irow, row := range field.rows {
		for icol, col := range row {
			if col == '0' {
				starts = append(starts, TrailStart{irow, icol})
			}
		}
	}

	sum := 0
	for _, start := range starts {
		u.Println3(start)
		finishes, _ := field.findDistinctFinished(start.row, start.col)
		sum += len(finishes.items)
		u.Println3(finishes)
		u.Println3("========")
	}

	u.Println("Day 10, Part 1:", sum)

	return sum
}

func SolvePart2(fileName string) int {

	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)

	field := Field{rows: make([][]rune, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		field.rows = append(field.rows, []rune(line))
		u.Println3(line)
	}

	starts := make([]TrailStart, 0)
	for irow, row := range field.rows {
		for icol, col := range row {
			if col == '0' {
				starts = append(starts, TrailStart{irow, icol})
			}
		}
	}

	sum := 0
	for _, start := range starts {
		u.Println3(start)
		finishes, _ := field.findAllFinishes(start.row, start.col)
		sum += len(finishes)
		u.Println3(finishes)
		u.Println3("========")
	}

	u.Println("Day 10, Part 2:", sum)

	return sum
}
