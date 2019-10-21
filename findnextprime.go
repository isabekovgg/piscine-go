package piscine

// import "fmt"

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return FindNextPrime(nb+1)
			}
		}
	}
	return nb
}

// func main() {
// 	fmt.Println(FindNextPrime(14))
// 	fmt.Println(FindNextPrime(7))
// }
