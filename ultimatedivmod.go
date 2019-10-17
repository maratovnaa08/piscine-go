package piscine

//UltimateDivMod 2323
func UltimateDivMod(a *int, b *int) {
	d := *a
	c := *b
	*a = d / c
	*b = d % c
}
