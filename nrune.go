package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		t := []rune(s)
		return t[n-1]
	} else {
		return 0
	}
}

func len(s string) int {
	y := 0
	for range s {
		y++
	}
	return y
}

// func main() {
// 	z01.PrintRune(NRune("Hello!", 3))
// 	z01.PrintRune(NRune("Salut!", 2))
// 	z01.PrintRune(NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }
