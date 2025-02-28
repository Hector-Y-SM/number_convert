package main

import (
	"convertir/enteros" //para poder usar funciones de un modulo debe iniciar en mayuscula la funcion
	"convertir/romanos"
	"fmt"
)

func main() {
	test := "MMMCMXCIX"

	if romanos.Sintaxis(test) && romanos.TripleRestriccion(test) {
		fmt.Print(test, " = ", romanos.RomanToInt(test), "\n")
	} else {
		fmt.Print("numero con sintaxis no valida")
	}

	test2 := 3999
	if test2 >= 4000 {
		fmt.Print("Numero fuera de rango")
	} else {
		fmt.Print(test2, " = ", enteros.EnteroARomano(test2), "\n")
	}

}
