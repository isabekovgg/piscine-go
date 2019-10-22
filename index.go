package piscine

// import (
// 	"fmt"
// )

func Index(s string, toFind string) int {
	var x int
	str := []rune(s)
	sub := []rune(toFind)
	y := 0
	z := 0
	for range s {
		y++
	}
	for range toFind {
		z++
	}
	if y > 0 && z > 0 {
		for i := 0; i <= y-1; i++ {
			if str[i] == sub[0] {
				x = i
				break
			} else {
				x = -1
			}
		}
	} else {
		x = -1
	}
	return x
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
