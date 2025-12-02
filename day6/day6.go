package day6

import (
	"bufio"
	"os"
	"slices"
	"time"

	u "github.com/bennyharvey/aoc2024/utils"
)

type Coords struct {
	row int
	col int
}

type Direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Guard struct {
	pos Coords
	dir Direction
}

func (g *Guard) print() {
	u.Printf1("Guard is at position %+v, facing ", g.pos)
	if g.dir == UP {
		u.Println1("UP")
	}
	if g.dir == DOWN {
		u.Println1("DOWN")
	}
	if g.dir == LEFT {
		u.Println1("LEFT")
	}
	if g.dir == RIGHT {
		u.Println1("RIGHT")
	}
}

func SpawnGuard(startingPosition Coords, startingDir Direction) *Guard {
	return &Guard{pos: startingPosition, dir: startingDir}
}

type Game struct {
	field   [][]rune
	guard   Guard
	over    bool
	endless bool
	visited map[int]map[int][]Direction
}

func (g *Guard) move() {
	if g.dir == UP {
		g.pos.row -= 1
	}
	if g.dir == DOWN {
		g.pos.row += 1
	}
	if g.dir == LEFT {
		g.pos.col -= 1
	}
	if g.dir == RIGHT {
		g.pos.col += 1
	}
}

func NewGame(size int) *Game {
	field := make([][]rune, size)
	for i := 0; i < size; i++ {
		field[i] = make([]rune, size)
	}
	visited := make(map[int]map[int][]Direction)
	return &Game{field: field, over: false, endless: false, visited: visited}
}

func (g *Game) copy() *Game {
	field := make([][]rune, len(g.field))
	for i, row := range g.field {
		field[i] = make([]rune, len(row))
		copy(field[i], row)
	}
	guard := g.guard
	visited := make(map[int]map[int][]Direction)
	return &Game{field: field, guard: guard, over: false, endless: false, visited: visited}
}

func (g *Game) printField() {
	u.Println1("Game Field:")
	u.Printf1(" ")
	for x := 0; x < len(g.field); x++ {
		u.Printf1("%d", x)
	}
	u.Printf1("\n")
	for row := 0; row < len(g.field); row++ {
		u.Printf1("%d", row)
		for col := 0; col < len(g.field[row]); col++ {
			u.Printf1("%c", g.field[row][col])
		}
		u.Printf1("\n")
	}
}

func (game *Game) appendVisit() {
	_, exists := game.visited[game.guard.pos.row]
	if !exists {
		game.visited[game.guard.pos.row] = make(map[int][]Direction)
	}
	game.visited[game.guard.pos.row][game.guard.pos.col] = append(game.visited[game.guard.pos.row][game.guard.pos.col], game.guard.dir)
}

func (game *Game) alreadyVisited() bool {
	return slices.Contains(game.visited[game.guard.pos.row][game.guard.pos.col], game.guard.dir)
}

func (g *Game) print() {
	// u.ClearScreen()
	// g.printField()
	// u.Println1("over:", g.over, "endless:", g.endless)
	// g.guard.print()
	// u.Println1("visited", g.visited)
}

func (game *Game) nextMovePossible() bool {
	if game.guard.dir == UP {
		if game.guard.pos.row == 0 {
			return false
		}
	}
	if game.guard.dir == LEFT {
		if game.guard.pos.col == 0 {
			return false
		}
	}
	if game.guard.dir == DOWN {
		if game.guard.pos.row >= len(game.field)-1 {
			return false
		}
	}
	if game.guard.dir == RIGHT {
		if game.guard.pos.col >= len(game.field)-1 {
			return false
		}
	}
	return true
}

func (game *Game) tick(delay int) {
	if game.alreadyVisited() {
		game.over = true
		game.endless = true
		return
	}
	game.appendVisit()
	game.field[game.guard.pos.row][game.guard.pos.col] = 'X'
	game.guard.move()
	if !game.nextMovePossible() {
		game.field[game.guard.pos.row][game.guard.pos.col] = 'X'
		game.over = true
		return
	}
	if game.guard.dir == UP {
		if game.field[game.guard.pos.row-1][game.guard.pos.col] == '#' {
			game.guard.dir = RIGHT
			game.field[game.guard.pos.row][game.guard.pos.col] = '>'
		} else {
			game.field[game.guard.pos.row][game.guard.pos.col] = '^'
		}
	}
	if game.guard.dir == RIGHT {
		if game.field[game.guard.pos.row][game.guard.pos.col+1] == '#' {
			game.guard.dir = DOWN
			game.field[game.guard.pos.row][game.guard.pos.col] = 'v'
		} else {
			game.field[game.guard.pos.row][game.guard.pos.col] = '>'
		}
	}
	if game.guard.dir == DOWN {
		if game.field[game.guard.pos.row+1][game.guard.pos.col] == '#' {
			game.guard.dir = LEFT
			game.field[game.guard.pos.row][game.guard.pos.col] = '<'
		} else {
			game.field[game.guard.pos.row][game.guard.pos.col] = 'v'
		}
	}
	if game.guard.dir == LEFT {
		if game.field[game.guard.pos.row][game.guard.pos.col-1] == '#' {
			game.guard.dir = UP
			game.field[game.guard.pos.row][game.guard.pos.col] = '^'
		} else {
			game.field[game.guard.pos.row][game.guard.pos.col] = '<'
		}
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
	// game.print()
}

func SolvePart1(fileName string) {
	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)

	game := NewGame(10)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			game.field[row][col] = char
			if char == '^' {
				game.guard = *SpawnGuard(Coords{col: col, row: row}, UP)
			}
		}
		row++
	}

	for !game.over {
		game.tick(100)
	}

	markedPositions := 0
	for _, row := range game.field {
		for _, col := range row {
			if col == 'X' {
				markedPositions++
			}
		}
	}

	u.Println("Day 6, Part 1:", markedPositions)
}

func SolvePart2(fileName string, fieldSize int) int {
	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)

	game := NewGame(fieldSize)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			game.field[row][col] = char
			if char == '^' {
				game.guard = *SpawnGuard(Coords{col: col, row: row}, UP)
			}
		}
		row++
	}

	saveGame := game.copy()

	for !game.over {
		game.tick(0)
	}
	game.print()

	count := 0
	for irow, row := range saveGame.field {
		for icol, col := range row {
			if col == '#' {
				continue
			}
			newGame := saveGame.copy()
			newGame.field[irow][icol] = '#'
			for !newGame.over {
				newGame.tick(0)
			}
			if newGame.endless {
				newGame.print()
				count++
			}
		}
	}

	u.Println("Day 6, Part 2:", count)

	return count
}

// part 2 wrong answers
// 2135
//
