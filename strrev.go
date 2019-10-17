package piscine

//S PrinStr do something
func StrRev(s string) string {
	var abc string
	for _, b := range s {
		abc := string(b) + abc
	}
	return abc
}
