dict = {
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

def int_roman(num):
    values = sorted(dict.keys(), reverse=True)

    res = ""
    for value in values:
        simbolo = dict[value]
        while num >= value:
            res += simbolo
            num -= value

    return res