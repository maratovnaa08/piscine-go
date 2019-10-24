package piscine

func LastRune(s string) rune {
	str_rune := []rune(s)
	return str_rune[s-1]
}
