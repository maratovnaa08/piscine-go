package piscine

import "github.com/01-edu/z01.PrintRune"

func Raid1b(x,y int) {
	for i := 0; i < y; i++ ) { 
		for ( j := 0; j <x; j++) {
			if ( i = 0 && j == 0) || ( i = y-1 && j = x-1) {
				z01.PrintRune(47)
			}
			else { if ( i == o && j == x-1) || (i = y-1 && j = 0)
				z01.PrintRune(92)
			}
			else { if  ( i > 0 && i < x-1 && j > 0 && j < y-1)
				z01.PrintRune(32)
			}
			z01.PrintRune(42)
		}
	}
}




			

