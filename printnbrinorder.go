package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	ran := 0
	var new int
	arr := []rune{}

	if n == 0 {
		arr = append(arr, rune(n))
	}

	for n != 0 {
		new = n % 10
		n = n / 10
		arr = append(arr, rune(new))
	}
	for range arr {
		ran++
	}
	for ind1 := 0; ind1 < ran-1; ind1++ {
		for ind2 := ind1; ind2 < ran; ind2++ {
			if arr[ind1] > arr[ind2] {
				arr[ind1], arr[ind2] = arr[ind2], arr[ind1]
			}
		}
	}
	for length := 0; length < ran; length++ {
		z01.PrintRune(arr[length] + 48)
	}
}

// func main() {
// 	// PrintNbrInOrder(321)
// 	PrintNbrInOrder(321321)
// 	// PrintNbrInOrder(2376960573752928263)
// }
