package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	size := len(runes)
	runesRev := make([]rune, size)
	for i := size; i > 0; i-- {
		runesRev[size-i] = runes[i-1]
	}
	output = string(runesRev)
	return output
}
