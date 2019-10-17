package piscine

// import "fmt"

func BasicAtoi(s string) int {
	stringrune := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}

	for x := 0; x <= y-1; x++ {
		final = final*10 + int(stringrune[x]-'0')
	}
	return final
}

// func main() {
// 	s := "12345"
// 	s2 := "0000000012345"
// 	s3 := "000000"

// 	n := BasicAtoi(s)
// 	n2 := BasicAtoi(s2)
// 	n3 := BasicAtoi(s3)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// }
