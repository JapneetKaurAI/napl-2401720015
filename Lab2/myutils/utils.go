package myutils

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func CountVowels(s string) int {
	count := 0

	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			count++
		}
	}

	return count
}

func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

// Power calculates base raised to the power exponent.
func Power(a, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= a
	}

	return result
}
