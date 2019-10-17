package piscine

//S PrinStr do something
func StrLen(str string) int {
	a := 0
	for _, b := range str {
		if b == b {
			a++
		}

	}
	return a
}
