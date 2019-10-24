package piscine

func Index(s string, toFind string) int {
	s_rune := []rune(s)
	to_Find_rune := []rune(toFind)
	count_find := 0
	for range to_Find_rune {
		count_find = count_find + 1
	}
	for Index, str_s := range s_rune {
		if count_find > 0 && str_s == to_Find_rune[0] {
			return Index
		}
	}
	return -1
}
