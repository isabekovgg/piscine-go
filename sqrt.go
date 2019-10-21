package piscine

// import "fmt"

func Sqrt(nb int) int {
	var x int
	if nb > 1 {
		for i := nb - 1; i >= 2; i-- {
			if i == nb/i {
				x = i
			}
		}
	} else {
		return 0
	}
	return x
}

// func main() {
// 	arg1 := 4
// 	arg2 := 36
// 	fmt.Println(Sqrt(arg1))
// 	fmt.Println(Sqrt(arg2))

// }
