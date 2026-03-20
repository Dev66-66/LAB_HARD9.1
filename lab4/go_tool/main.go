package main

import (
	"fmt"
	"os"
	"strconv"
)

// Square calculates the square of a number.
func Square(n int) int {
	return n * n
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go_tool <number>")
		os.Exit(1)
	}

	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Invalid input: %s\n", arg)
		os.Exit(1)
	}

	fmt.Println(Square(n))
}
