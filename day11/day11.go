package day11

import (
	"bufio"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	u "github.com/bennyharvey/aoc2024/utils"
)

func digitCount(num int) (int, bool) {
	digitCount := 0
	for num > 0 {
		num /= 10
		digitCount++
	}
	return digitCount, digitCount%2 == 0
}

type Game struct {
	stones []int
}

func (game *Game) blinkTimes(times int) {
	for i := range times {
		start := time.Now()
		game.blink()
		elapsed := time.Since(start)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		// u.Println1(len(game.stones))
		// u.Println1(game.stones)
		u.Printf1("Blink #%d took %s, Memory: %vmb\n", i, elapsed, m.Sys/1024/1024)
	}
}

func solve(cache Cache, num int, depth int) int {
	if !contains(cache, num, depth) {
		result := 0
		digitCount, isEven := digitCount(num)
		if depth == 1 {
			return 1
		}
		if num == 0 {
			result = solve(cache, 1, depth-1)
		} else if isEven {
			result = 0
			result += solve(cache, num%(pow(10, (digitCount/2))), depth-1)
			result += solve(cache, num/pow(10, (digitCount/2)), depth-1)

		} else {
			result = solve(cache, num*2024, depth-1)
		}
		setCache(cache, num, depth, result)
		// cache[num][depth] = result
	}

	return cache[num][depth]
}

func (game *Game) blinkTimesMemoized(times int, cache Cache) int {
	count := 0
	for j := range game.stones {
		count += solve(cache, game.stones[j], times)
	}

	return count
}

func (game *Game) blinkMemoized() {
	for i := 0; i < len(game.stones); i++ {
		if game.stones[i] == 0 {
			game.stones[i] = 1
			continue
		}
		digitCount, isEven := digitCount(game.stones[i])
		if isEven {
			game.stones = append(game.stones, game.stones[i]%(pow(10, (digitCount/2))))
			game.stones[i] = game.stones[i] / pow(10, (digitCount/2))
			continue
		}
		game.stones[i] *= 2024
	}

	u.Println3(game.stones)
}
func (game *Game) blink() {
	len := len(game.stones)
	for i := 0; i < len; i++ {
		if game.stones[i] == 0 {
			game.stones[i] = 1
			continue
		}
		digitCount, isEven := digitCount(game.stones[i])
		if isEven {
			// game.stones = slices.Insert(game.stones, i+1, game.stones[i]%(pow(10, (digitCount/2))))
			game.stones = append(game.stones, game.stones[i]%(pow(10, (digitCount/2))))
			game.stones[i] = game.stones[i] / pow(10, (digitCount/2))
			// i++
			// len++
			continue
		}
		game.stones[i] *= 2024
	}

	// u.Println3(game.stones)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

type Cache map[int]map[int]int

func contains(cache Cache, num, depth int) bool {
	_, numExists := cache[num]
	if numExists {
		_, depthExists := cache[num][depth]
		return depthExists
	}
	return false
}

func setCache(cache Cache, num, depth, value int) {
	_, numExists := cache[num]
	if !numExists {
		cache[num] = make(map[int]int)
	}
	cache[num][depth] = value
}

func SolvePart1(fileName string) int {
	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	game := Game{make([]int, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		for _, num := range numbers {
			game.stones = append(game.stones, u.Must(strconv.Atoi(num)))
		}
		u.Println1(line)
	}

	// count := 0
	cache := make(Cache)
	count := game.blinkTimesMemoized(76, cache)
	// game.blinkTimes(50)

	u.Println(cache)
	u.Println("Day 11, Part 2:", count)

	return count
}

func SolvePart2(fileName string) int {
	file := u.Must(os.Open(fileName))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		u.Println3(line)
	}
	sum := 0
	u.Println("Day 11, Part 1:", sum)

	return sum
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
