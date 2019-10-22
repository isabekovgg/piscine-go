package piscine

// import "fmt"

func IsUpper(str string) bool {
	new := []rune(str)
	y := 0
	for range str {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] < 'A' || new[i] > 'Z' {
			return false
		}
	}
	return true
}
