from roman.convert import dict_num

def limit(input):
	count = 1
	for i in range(1, len(input)) :
		if input[i] == input[i-1] :
			count+= 1 
			if count > 3 :
				return False
		else:
			count = 1
	return True

def sintaxis(input_str):
    for i in range(len(input_str) - 1):
        current_char = input_str[i]
        next_char = input_str[i + 1]

        current_value = dict_num[current_char]
        next_value = dict_num[next_char]

        if current_value < next_value:
            if not (
                (current_char == "I" and next_char in ("V", "X")) or
                (current_char == "X" and next_char in ("L", "C")) or
                (current_char == "C" and next_char in ("D", "M"))
            ):
                print("esto no se puede:", current_char + next_char)
                return False
    return True