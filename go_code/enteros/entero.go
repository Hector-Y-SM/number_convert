package enteros

import (
	"sort"
)

func EnteroARomano(numero int) string {
	romanos := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	// extraer y ordenar las claves del mapa de mayor a menor
	values := make([]int, 0, len(romanos))
	for value := range romanos {
		values = append(values, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	res := ""
	for _, value := range values {
		simbolo := romanos[value]
		// repetir el símbolo mientras el número sea mayor o igual al value
		for numero >= value {
			res += simbolo
			numero -= value
		}
	}

	return res
}
