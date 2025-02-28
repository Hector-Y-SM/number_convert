from roman.convert import roman_to_int
from roman.validation import limit, sintaxis 
from int_roman import int_roman

test = "X"
if sintaxis(test) and limit(test):
	print(test, " = ", roman_to_int(test), "\n")
else :
	print("numero con sintaxis no valida")

test2 = 3999
if test2 >= 4000:
	print("Numero fuera de rango")
else :
	print(test2, " = ", int_roman(test2), "\n")

