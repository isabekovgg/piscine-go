package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	len := 0
	arg := os.Args
	for range arg {
		len++
	}
	for x := len - 1; x > 0; x-- {
		mas := []rune(arg[x])
		for z := range mas {
			z01.PrintRune(mas[z])
		}
		z01.PrintRune(10)

	}
}
