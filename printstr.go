package main

import "github.com/01-edu/z01"

//S PrinStr do something
func PrintStr(str string) {
	for _, b := range str {
		z01.PrintRune(str)
	}
}
