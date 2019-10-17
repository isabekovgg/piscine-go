package main

// import "fmt"

func StrRev(s string) string {
	var reverse string
	for x := len(s)-1; x >= 0; x-- {
		reverse += string(s[x])
	}
	return reverse
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
