package piscine

func Index(s string, toFind string) int {
	//s_rune := []rune(s)
	toFind_rune := []rune(toFind)
	S_rune := []rune(s)
	count_find := 0
	for range toFind_rune {
		count_find = count_find + 1
	}
	count_s := 0
	for range S_rune {
		count_s++
	}
	for index := 0; index < count_s; index++ {
		if count_find > 0 {
			if string(S_rune[index]) == string(toFind_rune[0]) {
				return index
			} else {
				return -1
			}	
		} else {
			return -1
		}
	}
	return -1
}
