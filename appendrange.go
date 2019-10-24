package piscine

// import "fmt"

func AppendRange(min, max int) []int {

	var ans []int

	for i := min; i < max; i++ {
		ans = append(ans, i)
	}
	return ans
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 5))
// }
