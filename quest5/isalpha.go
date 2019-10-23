package piscine

// import "fmt"

func IsAlpha(str string) bool {
	new := []rune(str)
	y := 0
	for range str {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] < '0' || (new[i] > '9' && new[i] < 'A') || (new[i] > 'Z' && new[i] < 'a') || new[i] > 'z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))

// }
