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
			for i := range mas {
				z01.PrintRune(mas[i])
			}
			z01.PrintRune(10)
		}
	}
}
