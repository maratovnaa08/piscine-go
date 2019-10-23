package piscine

func intToDigitss(s string) (digits []byte) {
	count := 0
	count_d := 0
	for range s {
		count = count + 1
	}
	if count >= 1 {
		for i := 0; i <= count-1; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				digits = append(digits, s[i])
			}
			for range digits {
				count_d = count_d + 1
			}
			if s[i] == '-' && count_d == 0 {
				digits = append(digits, s[i])
			}
		}
	}
	return
}

func TrimAtoi(s string) int {
	arr_byte := intToDigitss(s)
	dec := 0
	signNumb := 0
	index := 0
	var signV rune
	for i, j := range arr_byte {
		ed := 0
		if j == '-' || j == '+' {
			index = i
			signNumb++
			if j == '-' {
				signV = '-'
			}
			if index > 0 {
				return 0
			}
		} else if j < '0' || j > '9' || signNumb > 1 {
			return 0
		}
		for i := '1'; byte(i) <= j; i++ {
			ed = ed + 1
		}
		dec = dec*10 + ed
	}
	if signV == '-' {
		return dec * -1
	} else {
		return dec
	}
}
