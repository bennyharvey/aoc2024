package day24

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	u "github.com/bennyharvey/aoc2024/utils"
)

type Gate struct {
	leftWire         string
	op               string
	rightWire        string
	resWire          string
	left, right, res bool
}

func SolvePart1(fileName string) int {
	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)

	toProcess, valueMap := parseGates(scanner)
	processed := processGates(toProcess, valueMap)
	res := getZBits(processed)

	fmt.Println(strconv.FormatInt(int64(res), 2))
	u.Println("Day 24, Part 1:", res)
	errors.New("123")

	return 0
}

func SolvePart2(fileName string) int {

	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)

	toProcess, valueMap := parseGates(scanner)
	u.Print(valueMap)
	x, y := parseXY(valueMap)
	u.Print(strconv.FormatInt(int64(x), 2), x, strconv.FormatInt(int64(y), 2), y)
	processed := processGates(toProcess, valueMap)
	res := getZBits(processed)
	u.Print(valueMap)

	u.Print(x&y == res)
	printGates(processed)
	u.Print(strconv.FormatInt(int64(res), 2))
	u.Println("Day 24, Part 2:", res)

	// part 1 - 48806532300520
	return 0
}

func remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func printGates(gates []Gate) {
	for _, g := range gates {
		u.Printf("%v %s %v = %v", g.leftWire, g.op, g.rightWire, g.resWire)
		u.Printf("	%d %s %d = %d\n", i(g.left), g.op, i(g.right), i(g.res))
	}
	u.Print("")
}

func parseGates(s *bufio.Scanner) ([]Gate, map[string]bool) {
	valueMap := make(map[string]bool)
	toProcess := make([]Gate, 0)

	firstSection := true

	for s.Scan() {
		line := s.Text()
		// u.Println3(line)

		if line == "" {
			firstSection = false
			continue
		}

		if firstSection {
			p := strings.Split(line, ": ")
			valueMap[p[0]] = u.Must(strconv.Atoi(p[1])) == 1
		} else {
			p := strings.Split(line, " ")
			toProcess = append(toProcess, Gate{
				leftWire:  p[0],
				op:        p[1],
				rightWire: p[2],
				resWire:   p[4],
			})
		}
	}

	return toProcess, valueMap
}

func parseXY(values map[string]bool) (uint, uint) {
	size := len(values) / 2
	tx := make([]bool, size)
	ty := make([]bool, size)
	for k, v := range values {
		if k[0] == 'x' {
			i := u.Must(strconv.Atoi(strings.Split(k, "x")[1]))
			tx[size-i-1] = v
		}
		if k[0] == 'y' {
			i := u.Must(strconv.Atoi(strings.Split(k, "y")[1]))
			ty[size-i-1] = v
		}
	}
	x := toBin(tx)
	y := toBin(ty)
	return x, y
}

func processGates(toProcess []Gate, valueMap map[string]bool) []Gate {
	processed := make([]Gate, 0)
	prevLen := len(toProcess)
	for {
		for i := 0; i < len(toProcess); i++ {
			gate := toProcess[i]
			left, leftKnown := valueMap[gate.leftWire]
			right, rightKnown := valueMap[gate.rightWire]
			if leftKnown && rightKnown {
				gate.left = left
				gate.right = right
				switch gate.op {
				case "AND":
					gate.res = gate.left && gate.right
				case "OR":
					gate.res = gate.left || gate.right
				case "XOR":
					gate.res = gate.left != gate.right
				}
				valueMap[gate.resWire] = gate.res
				processed = append(processed, gate)
				toProcess = remove(toProcess, i)
				i++
			}
		}

		if len(toProcess) == 0 {
			break
		}

		if prevLen == len(toProcess) {
			u.Println3(valueMap)
			u.Println3("==============")
			printGates(toProcess)
			u.Println3("==============")
			log.Fatal("process queue did not change")
		}
		prevLen = len(toProcess)
	}

	return processed
}

func getZBits(gates []Gate) uint {
	zvalues := make(map[int]bool, 0)
	keys := make([]int, 0)
	for _, gate := range gates {
		if gate.resWire[0] == 'z' {
			d := strings.Split(gate.resWire, "z")
			key := u.Must(strconv.Atoi(d[1]))
			zvalues[key] = gate.res
			keys = append(keys, key)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var res uint
	for _, i := range keys {
		if zvalues[i] {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}
	return res
}

func toBin(s []bool) uint {
	var res uint
	for _, i := range s {
		if i {
			res = (res << 1) | 1
		} else {
			res = res << 1
		}
	}
	return res
}
