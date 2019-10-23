package piscine

// import "github.com/01-edu/z01"

func LastRune(s string) rune {
	y := 0
	for range s {
		y++
	}
	t := []rune(s)
	return t[y-1]
}

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }
