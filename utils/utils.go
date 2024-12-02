package utils

import "log"

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
