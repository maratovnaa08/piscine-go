package piscine

func NRune(s string, n int) rune {
	s_rune := []rune(s)
	count := 0
	for range s_rune {
		count = count + 1
	}
	if n <= count-1 {
		return s_rune[n-1]
	}
	return 0
}
