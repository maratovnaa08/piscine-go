package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20000 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb == 1 {
		return 1
	} else {
		return RecursiveFactorial(nb-1) * nb
	}
	return 0
}
