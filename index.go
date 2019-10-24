package piscine

func Index(s string, toFind string) int {
	//s_rune := []rune(s)
	toFind_rune := []rune(toFind)
	count_find := 0
	for range toFind_rune {
		count_find = count_find + 1
	}
	count_s := 0
	for range s {
		count_s++
	}
	for index := 0; index < count_s; index++ {
		if count_find > 0 && string(s[index]) == string(toFind[0]) {
			return index
		}
	}
	return -1
}
