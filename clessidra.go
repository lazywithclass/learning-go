package main

import (
	"fmt"
)

func main() {
	clessidra(1, 0)
	clessidra(3, 0)
	clessidra(9, 0)
}

func clessidra(n, offset int) {
	if n == 1 {
		printBlanks(offset)
		printStars(1)
		return
	}

	printBlanks(offset)
	printStars((2 * n) - 1)
	clessidra(n - 1, offset + 1)
	printBlanks(offset)
	printStars((2 * n) - 1)
}

func printBlanks(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func printStars(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
}
