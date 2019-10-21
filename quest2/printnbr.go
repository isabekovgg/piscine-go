package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	x := 0

	for i := n; i < 1 && i > -1 {
		x++
		i = i/10
	}
	for a := x; a >= 0; a-- {
		z01.PrintRune('n % (a+1)')
	}

}

func main() {
	z01.PrintNbr(-123)
	z01.PrintNbr(0)
	z01.PrintNbr(123)
}
