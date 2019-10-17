package piscine

import "fmt"

func Swap(a *int, b *int) {
	x := *a
	y := *b
	*b = x
	*a = y
}

// func main() {
// 	a := 0
// 	b := 1
// 	Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
