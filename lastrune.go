package piscine

func LastRune(s string) rune {
	a := 0
	for _, char := range s {
		if char == char {
			a++
		}
	}
	y := []rune(s)
	return y[a-1]
}
