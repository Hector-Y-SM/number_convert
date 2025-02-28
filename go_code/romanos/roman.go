package romanos

func RomanToInt(input string) int {
	total := 0
	prevValue := 0

	for i := len(input) - 1; i >= 0; i-- {
		currentChar := string(input[i])
		currentValue := romanValues[currentChar]
		//fmt.Print(currentValue, " -> ", i, "\n")

		if currentValue < prevValue {
			total -= currentValue // si un numero menor esta antes que uno mayor se resta
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}
	return total
}
