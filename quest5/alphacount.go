package piscine

// import "fmt"

func AlphaCount(str string) int {

	s := []rune(str)
	y := slen(str)
	count := 0
	for i := 0; i < y; i++ {
		if s[i] < 91 || s[i] > 96 {
			if s[i] > 'A' && s[i] < 'z' {
				count++
			}
		}
	}
	return count
}

func slen(s string) int {
	y := 0
	for range s {
		y++
	}
	return y
}

// func main() {
// 	str := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(str)
// 	fmt.Println(nb)
// }
