package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for x := '1'; x <= '8'; x++ {
			for y := '2'; y <= '9'; y++ {
				if i < x && x < y {
					z01.PrintRune(i)
					z01.PrintRune(x)
					z01.PrintRune(y)
					if i != '7' && x != '8' && y != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb()
}
