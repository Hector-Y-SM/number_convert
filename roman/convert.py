dict_num = {
    "I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}


def roman_to_int(input_str):
    total = 0
    prev_value = 0

    for i in range(len(input_str) - 1, -1, -1):  
        current_char = input_str[i]
        current_value = dict_num[current_char]
        # print(current_value, "->", i)

        if current_value < prev_value:
            total -= current_value  
        else:
            total += current_value

        prev_value = current_value

    return total