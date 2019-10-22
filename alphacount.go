package piscine

func AlphaCount(str string) int {
	counter := 0
	for range str {
		counter++
	}
	letter := []rune(str)
	for i := 0; i < counter-1; i++ {
		if (letter[i] >= 'a' || letter[i] <= 'z') && (letter[i] >= 'A' || letter[i] <= 'Z') {

			letter++
		}
		return (AlphaCount(letter))
	}
}
