package day6

import (
	"testing"
)

func TestDay6(t *testing.T) {
	if SolvePart2("day6_sample.txt", 10) != 6 {
		t.Fatal("NO")
	}
	if SolvePart2("tests/test1.txt", 10) != 6 {
		t.Fatal("NO")
	}
}
