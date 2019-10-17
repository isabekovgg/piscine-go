package piscine

import "fmt"

func StrLen(str string) int {
	y := 0
	for range str {
		y++
	}
	return y
}

// func main() {
// 	str := "Hello World!"
// 	nb := StrLen(str)
// 	fmt.Println(nb)
// }
