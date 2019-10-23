package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := os.Args[0]
	for i := range arg {
		z01.PrintRune(rune(arg[i]))
	}
}
