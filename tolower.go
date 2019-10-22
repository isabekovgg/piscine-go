package piscine

func ToLower(s string) string {
	y := 0
	new := []rune(s)
	for range s {
		y++
	}
	for i := 0; i < y; i++ {
		if new[i] >= 'A' && new[i] <= 'Z' {
			new[i] = (new[i] + 32)
		}
	}
	str := string(new)
	return str
}
