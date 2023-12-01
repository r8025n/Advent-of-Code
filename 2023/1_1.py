def get_value_from_line(line):
    value = ''

    for char in line:
        if char.isnumeric():
            value += char
    final_value = value[0] + value[-1]
    return int(final_value)


input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

output = 0

for line in input_lines:
    output += get_value_from_line(line)

print(output)
