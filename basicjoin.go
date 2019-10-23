package piscine

// import "fmt"

func BasicJoin(strs []string) string {
	var result string
	i := 0
	for range strs {
		result = result + strs[i]
		i++
	}
	return result
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(BasicJoin(toConcat))
// }
