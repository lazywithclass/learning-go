package main

import (
	"fmt"
)

func main() {
	clessidra(9, 0)
}

func clessidra(n, offset int) {
	if n == 0 {
		return
	}

	if n == 1 {
		printBlanks(offset)
		printStars((2 * n) - 1)
		clessidra(n - 1, offset + 1)
	} else {
		printBlanks(offset)
		printStars((2 * n) - 1)
		clessidra(n - 1, offset + 1)
		printBlanks(offset)
		printStars((2 * n) - 1)
	}
}

func printBlanks(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func printStars(n int) {
	if n == 0 {
		n = 1
	}
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
}
