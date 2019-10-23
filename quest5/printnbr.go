package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	i := n
	if n < 0 {
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	if i < 1 {
		z01.PrintRune{}
	}

}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
}
