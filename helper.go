package main

import (
	"fmt"
	"strconv"
)

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid number: %s", s))
	}

	return n
}
