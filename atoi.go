package main

import "fmt"

func Atoi(s string) int {
	stringrune := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}
	if int(stringrune[0] == '-' {
		for x := 1; x <= y-1; x++ {
			if int(stringrune[x]-'0') >=0  && int(stringrune[x]-'0') < 10 {
					final = final*10 + int(stringrune[x]-'0')		
				} else {
					return 0
			}
		}
	} else {
		for x := 0; x <= y-1; x++ {
			if int(stringrune[0]-'0') >=0  && int(stringrune[x]-'0') < 10 {
					final = final*10 + int(stringrune[x]-'0')		
				} else {
					return 0
			}
		}
	}
	return final
}


func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}