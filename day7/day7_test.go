package day7

import (
	"testing"
)

func TestDay6(t *testing.T) {
	if SolvePart1("tests/test1.txt") != 100 {
		t.Fatal("NO")
	}
	if SolvePart1("tests/test2.txt") != 100 {
		t.Fatal("NO")
	}
	if SolvePart1("tests/test3.txt") != 1 {
		t.Fatal("NO")
	}
}
