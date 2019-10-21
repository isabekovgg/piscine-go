package piscine

// import "fmt"

func IterativePower(nb int, power int) int {
	if power > 0 {
		result := 1
		for x := 1; x <= power; x++ {
			result = result*nb
		}
		return result
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}

// func main() {
// 	arg1 := 4
// 	arg2 := 3
// 	fmt.Println(IterativePower(arg1, arg2))
// }
