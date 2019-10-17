package piscine

//Swap 2424
func Swap(a *int, b *int) {
	d := *a
	c := *b
	*a = c
	*b = d
}
