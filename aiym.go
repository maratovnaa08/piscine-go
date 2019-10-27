package main

import "github.com/01-edu/z01"

func PrC() {
	for comb1 := '0'; comb1 <= '9'; comb1++ {
		for comb2 := '0'; comb2 <= '9'; comb2++ {
			z01.PrintRune(comb1)
			z01.PrintRune(comb2)
			if comb1 < '9' || comb2 < '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	PrC()
}
