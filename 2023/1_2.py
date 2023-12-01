number_map = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4', 
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9' 
}

def get_value_from_line(line):
    number_string = ''

    for index, char in enumerate(line):
        if char.isnumeric():
            number_string += char
        else:
            for key, val in number_map.items():
                if line[index : index + len(key)] == key:
                    number_string += val

    output = number_string[0] + number_string[-1]
    return int(output)

input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

output = 0



for line in input_lines:
    output += get_value_from_line(line)

print(output)