package piscine

// import "fmt"

func TrimAtoi(s string) int {
	final := 0
	sign := 1
	for _, com := range s {
		if com >= '0' && com <= '9' {
			final = final*10 + int(com-48)
		} else if com == '-' && final == 0 {
			sign = -1
		}
	}
	return final * sign
}

// func main() {
// 	s := "12345"
// 	s2 := "str123ing45"
// 	s3 := "012 345"
// 	s4 := "Hello World!"
// 	s5 := "sd+x1fa2W3s4"
// 	s6 := "sd-x1fa2W3s4"
// 	s7 := "sdx1-fa2W3s4"

// 	n := TrimAtoi(s)
// 	n2 := TrimAtoi(s2)
// 	n3 := TrimAtoi(s3)
// 	n4 := TrimAtoi(s4)
// 	n5 := TrimAtoi(s5)
// 	n6 := TrimAtoi(s6)
// 	n7 := TrimAtoi(s7)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// 	fmt.Println(n4)
// 	fmt.Println(n5)
// 	fmt.Println(n6)
// 	fmt.Println(n7)
// }
