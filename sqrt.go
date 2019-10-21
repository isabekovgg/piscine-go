package piscine

// import "fmt"

func Sqrt(nb int) int {
	x := 0
	if nb >= 1 {
		for i := nb; i >= 1; i-- {
			if i * i == nb {
				x = i
			}
		}
	} else {
		return 0
	}
	return x
}

// func main() {
// 	arg1 := 10
// 	arg2 := 5
// 	fmt.Println(Sqrt(arg1))
// 	fmt.Println(Sqrt(arg2))

// }
