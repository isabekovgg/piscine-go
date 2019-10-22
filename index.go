package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	sub := []rune(toFind)
	len1 := 0
	len2 := 0
	for index := range sub {
		index = index
		len2++
	}
	if len2 == 0 {
		return 0
	}
	for index := range str {
		index = index
		len1++
	}
	for index, letter := range str {
		if letter == sub[0] && len1 >= len2+index-1 {
			m := 1
			for i := 1; i < len2; i++ {
				if sub[i] == str[index+i] {
					m++
				}
			}
			if m == len2 {
				return index
			}
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
