package piscine

//NRune returns the nth rune of a string
func NRune(s string, n int) rune {
	l := 0
	for i := range s {
		if i == i {
			l++
		}
	}
	if n-1 >= 0 && l >= n {
		a := []rune(s)
		d := n - 1
		return a[d]
	}
	return 0

}
