package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	} else {
		return RecursivePower(power-1) * nb
	}
	return 0
}
