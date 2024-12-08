package day8

import (
	"bufio"
	"os"

	u "github.com/bennyharvey/aoc2024/utils"
)

type Node struct {
	row int
	col int
}

type Field struct {
	rows  [][]rune
	nodes map[rune][]Node
}

func NewField() *Field {
	return &Field{nodes: make(map[rune][]Node)}
}

func (f *Field) setAntinodes(n1 Node, n2 Node) {
	f.setAntinode1(n1, n2)
	f.setAntinode2(n1, n2)
}

func (f *Field) setAntinode1(n1 Node, n2 Node) {
	if n1 == n2 {
		return
	}
	f.rows[n1.row][n1.col] = '#'
	f.rows[n2.row][n2.col] = '#'

	dcol1 := n1.col - (n1.col-n2.col)*2
	drow1 := n1.row + (n2.row-n1.row)*2

	if drow1 >= len(f.rows) || drow1 < 0 {
		return
	}
	if dcol1 >= len(f.rows[drow1]) || dcol1 < 0 {
		return
	}
	f.rows[drow1][dcol1] = '#'

	f.setAntinode1(n2, Node{col: dcol1, row: drow1})
}

func (f *Field) setAntinode2(n1 Node, n2 Node) {
	if n1 == n2 {
		return
	}
	f.rows[n1.row][n1.col] = '#'
	f.rows[n2.row][n2.col] = '#'

	dcol2 := n1.col + (n1.col - n2.col)
	drow2 := n1.row - (n2.row - n1.row)

	if drow2 >= len(f.rows) || drow2 < 0 {
		return
	}
	if dcol2 >= len(f.rows[drow2]) || dcol2 < 0 {
		return
	}
	f.rows[drow2][dcol2] = '#'

	f.setAntinode2(n2, Node{col: dcol2, row: drow2})
}

func (f *Field) print() {
	for _, row := range f.rows {
		u.Printf1("%+c\n", row)
	}
	for char, node := range f.nodes {
		u.Printf1("%c, %+v\n", char, node)
	}
}

func SolvePart1(fileName string) int {

	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)
	field := NewField()

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		field.rows = append(field.rows, []rune(line))
		for col, char := range line {
			if char != '.' {
				field.nodes[char] = append(field.nodes[char], Node{col: col, row: row})
			}
		}
		row++
	}
	// field.print()
	for char, nodeType := range field.nodes {
		for _, node1 := range nodeType {
			for _, node2 := range field.nodes[char] {
				field.setAntinodes(node1, node2)
			}
		}
	}
	// field.print()

	locations := 0
	for _, row := range field.rows {
		for _, char := range row {
			if char == '#' {
				locations++
			}
		}
	}

	// for _, nodeType := range field.nodes {
	// 	if len(nodeType) > 1 {
	// 		locations += len(nodeType)
	// 	}
	// }

	// field.setAntinodpes(field.nodes['0'][2], field.nodes['0'][3])
	u.Println("Day 8, Part 2:", locations)
	return locations
}
