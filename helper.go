package main

import (
	"fmt"
	"strconv"
)

func getInt(s string) int {
	// v := strings.Trim(s, "\n")

	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("len s = %d\n", len(s))
		for _, ch := range s {
			fmt.Printf("-'%c'-", ch)
		}
		fmt.Println()
		panic(fmt.Sprintf("invalid number: '%s'", s))
	}

	return n
}
