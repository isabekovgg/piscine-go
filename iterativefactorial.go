package piscine

// import "fmt"

func IterativeFactorial(nb int) int {
	if nb > -20 && nb < 20{
		result := 1
		for x := 1; x <= nb; x++ {
			result = result * x
		}
		return result
	}
}

// func main() {
// 	arg := 3
// 	fmt.Println(IterativeFactorial(arg))
// }
