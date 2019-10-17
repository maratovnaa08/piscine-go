package piscine

//S PrinStr do something
func StrRev(s string) string {
	var a string
	for _, b := range s {
		a := string(b) + a
	}
	return a
}
