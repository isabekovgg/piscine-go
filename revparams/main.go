package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := os.Args

	for j := range arg {
		if j > 0 {
			mas := []rune(arg[j])
			i := 0
			for range mas {
				i++
			}
			for z := i - 1; z >= 0; z-- {
				z01.PrintRune(mas[z])
			}
			z01.PrintRune(10)
		}
	}
}
