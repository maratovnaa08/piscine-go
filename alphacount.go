package piscine

func AlphaCount(str string) int {
	counter := 0
	for range str {
		counter++
	}
	str_rune := []rune(str)
	for i := 0; i < counter-1; i++ {
		if (str_rune[i] >= 'a' || str_rune[i] <= 'z') && (str_rune[i] >= 'A' || str_rune[i] <= 'Z') {
			letter + 1
		}
		return letter
	}
}
