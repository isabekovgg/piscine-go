package piscine

// import "fmt"

func ConcatParams(args []string) string {
	var ans string
	y := 0
	for range args {
		y++
	}
	for i := 0; i < y; i++ {

		ans = ans + args[i]
		if i != y-1 {
			ans = ans + "\n"
		}
	}
	return ans
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
