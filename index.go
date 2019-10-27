package student

//Index behaves like the Index function
func Index(s string, toFind string) int {
	//b := []rune(s)
	k := 0
	d := 0
	for i := range toFind {
		if i == i {
			k++
		}
	}
	if k != 0 {
		c := []rune(toFind)
		for i, chr := range s {
			if i == i {
				d++
			}
			if c[0] == chr && c[0] != 10 {
				return i
			}
		}
		return -1
	}
	return 0
}
