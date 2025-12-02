package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

const DEBUG = true
const LOG_LEVEL = 3

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

func Printf(format string, a ...any) {
	if LOG_LEVEL > 0 {
		fmt.Printf(format, a...)
	}
}

func Print(a ...interface{}) {
	if LOG_LEVEL <= 0 {
		return
	}
	for _, i := range a {
		switch reflect.ValueOf(i).Kind() {
		case reflect.Map:
			fmt.Print(string(Must(json.MarshalIndent(i, "", "  "))), "\n")
		default:
			Println(i)
		}
		Println("===================================================")
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
	if LOG_LEVEL < 1 {
		return
	}
	for _, el := range slice {
		fmt.Printf("%+v\n", el)
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
