package utils

import (
	"fmt"
	"log"
)

const DEBUG = true
const LOG_LEVEL = 0

// Fatals on error, otherwise returns r
func Must[T any](r T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// Fatals on error
func Fie(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Printf1(format string, a ...any) {
	if LOG_LEVEL > 0 {
		fmt.Printf(format, a...)
	}
}

func Println1(a ...any) {
	if LOG_LEVEL > 0 {
		Println(a...)
	}
}

func Println2(a ...any) {
	if LOG_LEVEL > 1 {
		Println(a...)
	}
}

func Println3(a ...any) {
	if LOG_LEVEL > 2 {
		Println(a...)
	}
}

func Println(a ...any) {
	fmt.Println(a...)
}

func PPrintSlice[T any](slice []T) {
	for _, el := range slice {
		fmt.Printf("%+v\n", el)
	}
}
