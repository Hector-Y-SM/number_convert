package romanos

import "fmt"

var romanValues = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// funcion para delimitar el mismo caractar a no mas de 3 repeticiones
func TripleRestriccion(input string) bool {
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
			if count > 3 {
				return false
			}
		} else {
			count = 1
		}
	}

	return true
}

// funcion para decidir que puede restar y q no
func Sintaxis(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		currentChar := string(input[i])
		nextChar := string(input[i+1])

		currentValue := romanValues[currentChar]
		nextValue := romanValues[nextChar]

		if currentValue < nextValue {
			// cosas que SI se pueden restar
			if !(currentChar == "I" && (nextChar == "V" || nextChar == "X")) &&
				!(currentChar == "X" && (nextChar == "L" || nextChar == "C")) &&
				!(currentChar == "C" && (nextChar == "D" || nextChar == "M")) {
				fmt.Println("esto no se puede", currentChar+nextChar)
				return false
			}
		}
	}
	return true
}
