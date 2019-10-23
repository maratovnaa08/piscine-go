package piscine

func NRune(s string, n int) rune {
	str_rune := []rune(s)
	count := 0
	for range str_rune {
		count = count + 1
	}
	if n <= count-1 {
		return str_rune[n-1]
	}
	return 0
}
