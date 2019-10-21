package piscine

// import "fmt"

func UltimateDivMod(a *int, b *int) {

	var x int
	var y int
	x = *a / *b
	y = *a % *b
	*a = x
	*b = y

}

// func main() {
// 	a := 13
// 	b := 2
// 	UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
