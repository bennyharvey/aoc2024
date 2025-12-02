package day7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	u "github.com/bennyharvey/aoc2024/utils"
)

type Eval struct {
	expectedResult int64
	operands       []int64
}

func SolvePart1(fileName string) int64 {
	file := u.Must(os.Open(fileName))
	scanner := bufio.NewScanner(file)

	evals := make([]Eval, 0)
	for scanner.Scan() {
		eval := Eval{operands: make([]int64, 0)}
		line := scanner.Text()
		parts := strings.Split(line, ":")
		eval.expectedResult = u.Must(strconv.ParseInt(parts[0], 10, 64))
		parts = strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, part := range parts {
			eval.operands = append(eval.operands, u.Must(strconv.ParseInt(part, 10, 64)))
		}
		evals = append(evals, eval)
	}

	var sum int64
	sum = 0
	rline := ""

	for _, eval := range evals {
		u.Println(eval)
		slots := len(eval.operands) - 1
		carry := make([]string, 0)
		combinations := comb([]string{"+", "*"}, carry, 0, slots)

		// u.Println1(len(combinations))

		for _, operator := range combinations {
			var result int64
			result = 0
			line := ""
			for i := range eval.operands {
				if i == 0 {
					result = eval.operands[i]
					// line += strconv.Itoa(eval.operands[i])
					continue
				}
				if operator[i-1] == "+" {
					result += eval.operands[i]
					// line += " + " + strconv.Itoa(eval.operands[i])
				}
				if operator[i-1] == "*" {
					result *= eval.operands[i]
					// line += " * " + strconv.Itoa(eval.operands[i])
				}
			}
			if result == eval.expectedResult {
				u.Println3(eval, line, result)
				sum += eval.expectedResult
				// rline += " + " + strconv.Itoa(eval.expectedResult)
				break
			}
		}
	}

	u.Println("Day 7, Part 1:", sum, rline)
	return sum
}

func comb(arr []string, carry []string, step int, size int) [][]string {
	results := make([][]string, 0)
	if step == size {
		return [][]string{carry}
	}
	for i := range len(arr) {
		newCarry := append(carry, arr[i])
		results = append(results, comb(arr, newCarry, step+1, size)...)
	}
	return results
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// part 1 wrong results
// 1031636219777
// 285303239298899

// (await fetch("https://adventofcode.com/2024/day/7/input").then(resp=>resp.text()))
// .split('\n')
// .map(s=>s.split(':')
//   .reduce((n,nums)=>[eval(n),eval(`[${nums.trim().replaceAll(' ',',')}]`)]))
// .filter(s=>s)
// .reduce((sum,[n,nums])=>function f(n,nums,m,ns){
//   return 0<=n && Number.isInteger(n) && (
// 1==nums.length ? n==nums[0] :
// (m=n+'').endsWith(ns=nums.at(-1)+'')
// && f(eval(m.slice(0,-ns.length)),nums.slice(0,-1))
// || f(n/nums.at(-1), nums.slice(0,-1))
// || f(n-nums.at(-1), nums.slice(0,-1)))
//   }(n,nums)*n+sum,0)
