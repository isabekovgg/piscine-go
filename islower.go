package piscine

// import "fmt"

func IsLower(str string) bool {
	new := []rune(str)
	y := 0
	for range str {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] < 'a' || new[i] > 'z' {
			return false
		}
	}
	return true
}
