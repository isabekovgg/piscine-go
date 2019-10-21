package piscine

// import "fmt"

func StrRev(s string) string {
	var reverse string
	y := 0
	for range s {
		y++
	}

	for x := y - 1; x >= 0; x-- {
		reverse += string(s[x])
	}
	return reverse
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
