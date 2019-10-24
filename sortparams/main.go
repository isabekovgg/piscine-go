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
	for ind1 := 1; ind1 < len-1; ind1++ {
		for ind2 := ind1; ind2 < len; ind2++ {
			if arg[ind1] > arg[ind2] {
				arg[ind1], arg[ind2] = arg[ind2], arg[ind1]
			}
		}
	}
	for x := 1; x < len; x++ {
		mas := []rune(arg[x])
		for z := range mas {
			z01.PrintRune(mas[z])
		}
		z01.PrintRune(10)

	}
}
