
input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

for line in input_lines:
    line = line.rstrip('\n') + '.'

valid_numbers = []

def if_any_special_around(position_list):
    for position in position_list:
        if position[0] >= 0 and position[0] < len(input_lines)  and position[1] >= 0 and position[1] < len(input_lines[0]):
            if input_lines[position[0]][position[1]] not in  ['.', '\n'] and input_lines[position[0]][position[1]].isnumeric() == False:
                return True

    return False


if __name__=="__main__": 

    for i, line in enumerate(input_lines):
        number = ''
        possible_special_character_indices = []

        for j, char in enumerate(line):
            if char.isnumeric():
                possible_special_character_indices.extend([[i, j+1], [i, j-1], [i+1, j], [i-1, j], [i-1, j - 1], [i+1, j+1], [i+1, j-1], [i-1, j+1]])
                number += char
            else:
                if number != '':
                    if if_any_special_around(possible_special_character_indices):
                        valid_numbers.append(int(number))
                    number = ''
                    possible_special_character_indices = []

    # print(valid_numbers)
    print(sum(valid_numbers))                  




