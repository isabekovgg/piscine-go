package piscine

// import "fmt"

func BasicAtoi2(s string) int {
	stringrune := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}
	for x := 0; x <= y-1; x++ {
		if int(stringrune[x]-'0') >= 0 && int(stringrune[x]-'0') < 10 {
		final = final*10 + int(stringrune[x]-'0')
		} else {
			return 0
		}
	}
	return final
}

// func main() {
// 	s := "12345"
// 	s2 := "0000000012345"
// 	s3 := "012 345"
// 	s4 := "Hello World!"

// 	n := BasicAtoi2(s)
// 	n2 := BasicAtoi2(s2)
// 	n3 := BasicAtoi2(s3)
// 	n4 := BasicAtoi2(s4)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// 	fmt.Println(n4)

// }
