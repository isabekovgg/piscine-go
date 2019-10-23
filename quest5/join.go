package piscine

// import "fmt"

func Join(strs []string, sep string) string {
	var result string
	i := 0
	j := 0
	for range strs {
		j++
	}
	for range strs {
		if i != j-1 {
			result = result + strs[i] + sep
			i++
		} else {
			result = result + strs[i]
		}
	}
	return result
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
