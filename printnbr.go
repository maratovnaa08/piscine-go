package piscine

import (
	"github.com/01-edu/z01"
)

//PrintNbr do smth
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintN(n)
}

// PrintN get number
func PrintN(s int) {

	i := '0'

	if s == 0 {
		z01.PrintRune('0')
		return
	}
	for l := 1; l <= s%10; l++ {
		i++
	}
	for l := -1; l >= s%10; l-- {
		i++
	}
	if s/10 != 0 {
		PrintN(s / 10)
	}
	z01.PrintRune(i)
	return
}
