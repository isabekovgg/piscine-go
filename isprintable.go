package piscine

// import "fmt"

func IsPrintable(str string) bool {
	new := []rune(str)
	y := 0
	for range str {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] < ' ' {
			return false
		}
	}
	return true
}
