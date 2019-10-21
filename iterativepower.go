package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		x := 1
		for i := 1; i <= power; i++ {
			i = i * nb
			return x
		} else
		power = 0 {
			return 1
		}
	}
	return 0
}
