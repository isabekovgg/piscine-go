package piscine

// import "fmt"

func Compare(a, b string) int {
	var x int
	if a > b {
		x = 1
	}
	if a == b {
		x = 0
	}
	if a < b {
		x = -1
	}
	return x
}

// func main() {
// 	fmt.Println(Compare("Hello!", "Hello!"))
// 	fmt.Println(Compare("Salut!", "lut!"))
// 	fmt.Println(Compare("Ola!", "Ol"))
// }
