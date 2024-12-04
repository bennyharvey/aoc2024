package day4

import (
	"bufio"
	"os"

	u "github.com/bennyharvey/aoc2024/utils"
)

// Approach 1:
//   - Iterate all lines and save "X" locations to a sctruct
//   - Iterate all "X" locations and look in all 8 directions
//     to find "XMAS"

type Coords struct {
	x int
	y int
}

type Line struct {
	letters map[int]rune
}

func NewLine() *Line {
	return &Line{letters: make(map[int]rune, 0)}
}

type Field struct {
	lines map[int]Line
}

func NewField() *Field {
	return &Field{lines: make(map[int]Line, 0)}
}

func (f *Field) locate(x, y int) rune {
	_, xExists := f.lines[y]
	if !xExists {
		return '-'
	}
	_, yExists := f.lines[y].letters[x]
	if !yExists {
		return '-'
	}

	return f.lines[y].letters[x]
}

func SolvePart1(fileName string) int {
	Xs := make([]Coords, 0)
	field := NewField()

	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		field.lines[x] = *NewLine()
		for y, char := range line {
			if char == 'X' {
				Xs = append(Xs, Coords{x, y})
			}
			field.lines[x].letters[y] = char
		}
		u.Println3(line)
		x++
	}

	count := 0
	for _, X := range Xs {
		count += X.check(field)
	}

	// u.Println1(field)
	// u.PPrintSlice(Xs)

	u.Println("Day 4, Part 1: ", count)

	return count
}

func (X *Coords) check(f *Field) int {
	count := 0
	if f.locate(X.x, X.y+1) == 'M' {
		if f.locate(X.x, X.y+2) == 'A' {
			if f.locate(X.x, X.y+3) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x+1, X.y+1) == 'M' {
		if f.locate(X.x+2, X.y+2) == 'A' {
			if f.locate(X.x+3, X.y+3) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x+1, X.y) == 'M' {
		if f.locate(X.x+2, X.y) == 'A' {
			if f.locate(X.x+3, X.y) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x-1, X.y+1) == 'M' {
		if f.locate(X.x-2, X.y+2) == 'A' {
			if f.locate(X.x-3, X.y+3) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x, X.y-1) == 'M' {
		if f.locate(X.x, X.y-2) == 'A' {
			if f.locate(X.x, X.y-3) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x-1, X.y-1) == 'M' {
		if f.locate(X.x-2, X.y-2) == 'A' {
			if f.locate(X.x-3, X.y-3) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x-1, X.y) == 'M' {
		if f.locate(X.x-2, X.y) == 'A' {
			if f.locate(X.x-3, X.y) == 'S' {
				count++
			}
		}
	}
	if f.locate(X.x+1, X.y-1) == 'M' {
		if f.locate(X.x+2, X.y-2) == 'A' {
			if f.locate(X.x+3, X.y-3) == 'S' {
				count++
			}
		}
	}
	return count
}

type Direction int

const (
	DIR_RU = iota
	DIR_RD
	DIR_LU
	DIR_LD
)

func (M *Coords) countMasLeft(f *Field) ([]Direction, bool) {
	dirs := make([]Direction, 0)
	found := false
	if f.locate(M.x+1, M.y+1) == 'A' {
		if f.locate(M.x+2, M.y+2) == 'S' {
			dirs = append(dirs, DIR_RD)
			found = true
		}
	}
	if f.locate(M.x-1, M.y+1) == 'A' {
		if f.locate(M.x-2, M.y+2) == 'S' {
			dirs = append(dirs, DIR_LD)
			found = true
		}
	}
	if f.locate(M.x+1, M.y-1) == 'A' {
		if f.locate(M.x+2, M.y-2) == 'S' {
			dirs = append(dirs, DIR_RU)
			found = true
		}
	}
	if f.locate(M.x-1, M.y-1) == 'A' {
		if f.locate(M.x-2, M.y-2) == 'S' {
			dirs = append(dirs, DIR_LU)
			found = true
		}
	}
	return dirs, found
}

func (M *Coords) countMasRight(f *Field, dirs []Direction) int {
	count := 0
	for _, dir := range dirs {
		if dir == DIR_RD {
			x := M.x + 2
			y := M.y
			counter := f.locate(x, y)
			if counter == 'M' {
				if f.locate(x-1, y+1) == 'A' {
					if f.locate(x-2, y+2) == 'S' {
						count += 1
					}
				}
			}
			if counter == 'S' {
				if f.locate(x-1, y+1) == 'A' {
					if f.locate(x-2, y+2) == 'M' {
						count += 1
					}
				}
			}
		}
		if dir == DIR_LD {
			x := M.x - 2
			y := M.y
			counter := f.locate(x, y)
			if counter == 'M' {
				if f.locate(x+1, y+1) == 'A' {
					if f.locate(x+2, y+2) == 'S' {
						count += 1
					}
				}
			}
			if counter == 'S' {
				if f.locate(x+1, y+1) == 'A' {
					if f.locate(x+2, y+2) == 'M' {
						count += 1
					}
				}
			}
		}
		if dir == DIR_RU {
			x := M.x + 2
			y := M.y
			counter := f.locate(x, y)
			if counter == 'M' {
				if f.locate(x-1, y-1) == 'A' {
					if f.locate(x-2, y-2) == 'S' {
						count += 1
					}
				}
			}
			if counter == 'S' {
				if f.locate(x-1, y-1) == 'A' {
					if f.locate(x-2, y-2) == 'M' {
						count += 1
					}
				}
			}
		}
		if dir == DIR_LU {
			x := M.x - 2
			y := M.y
			counter := f.locate(x, y)
			if counter == 'M' {
				if f.locate(x+1, y-1) == 'A' {
					if f.locate(x+2, y-2) == 'S' {
						count += 1
					}
				}
			}
			if counter == 'S' {
				if f.locate(x+1, y-1) == 'A' {
					if f.locate(x+2, y-2) == 'M' {
						count += 1
					}
				}
			}
		}
	}
	return count
}

func (M *Coords) countMas(f *Field) int {
	count := 0
	dirs, found := M.countMasLeft(f)
	if found {
		count += M.countMasRight(f, dirs)
	}
	return count
}

func SolvePart2(fileName string) int {
	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)
	field := NewField()
	Ms := make([]Coords, 0)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		field.lines[y] = *NewLine()
		for x, char := range line {
			if char == 'M' {
				Ms = append(Ms, Coords{x, y})
			}
			field.lines[y].letters[x] = char
		}
		u.Println3(line)
		y++
	}

	count := 0
	for _, M := range Ms {
		count += M.countMas(field)
	}

	u.Println("Day 4, Part 2: ", count/2)

	return count
}
