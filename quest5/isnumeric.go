package piscine

// import "fmt"

func IsNumeric(str string) bool {
	new := []rune(str)
	y := 0
	for range str {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] < '0' || new[i] > '9' {
			return false
		}
	}
	return true
}
