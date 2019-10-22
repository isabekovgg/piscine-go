package piscine

// import "fmt"

func ToUpper(s string) string {
	y := 0
	new := []rune(s)
	for range s {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] >= 'a' && new[i] <= 'z' {
			new[i] = (new[i] - 32)
		}
	}
	str := string(new)
	return str
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
