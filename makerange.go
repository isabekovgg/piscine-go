package piscine

// import "fmt"

func MakeRange(min, max int) []int {
	if max > min {
		ans := make([]int, max-min)

		for i := 0; i < max-min; i++ {
			ans[i] = i + min
		}
		return ans
	} else {
		return nil
	}
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
