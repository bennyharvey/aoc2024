package day4

import "testing"

func Test1(t *testing.T) {
	files := []string{
		"tests/test1.txt",
		"tests/test2.txt",
		"tests/test3.txt",
		"tests/test4.txt",
		"tests/test5.txt",
		"tests/test6.txt",
		"tests/test7.txt",
		"tests/test8.txt",
		"tests/test9.txt",
	}
	for _, file := range files {
		if SolvePart1(file) != 1 {
			t.Fatal(file)
		}
	}

	if SolvePart1("tests/test11.txt") != 2 {
		t.Fatal("should be 2")
	}

	if SolvePart1("tests/test10.txt") != 18 {
		t.Fatal("sample should be 18")
	}
}

func Test2(t *testing.T) {
	files := []string{
		"tests/test12.txt",
	}
	for _, file := range files {
		if SolvePart2(file) != 1 {
			t.Fatal(file)
		}
	}
}
