package piscine

func Index(s string, toFind string) int {
	s_rune := []rune(s)
	toFind_rune := []rune(toFind)
	count_find := 0
	for range toFind_rune {
		count_find = count_find + 1
	}
	for index, str_s := range s_rune {
		if count_find > 0 && str_s == toFind_rune[0] {
			return index
		}
	}
	return -1
}
