
input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

for line in input_lines:
    line = line.rstrip('\n') + '.'

gear_attachment_map = {}
valid_numbers = []

def if_a_gear_attached(position_list):
    for position in position_list:
        if position[0] >= 0 and position[0] < len(input_lines)  and position[1] >= 0 and position[1] < len(input_lines[0]):
            if input_lines[position[0]][position[1]] not in  ['.', '\n'] and input_lines[position[0]][position[1]].isnumeric() == False:
                return True, str(position[0]) + ',' + str(position[1])

    return False, ''


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
                    gear_attached, gear_position = if_a_gear_attached(possible_special_character_indices)
                    if gear_attached:
                        if gear_position in gear_attachment_map:
                            gear_attachment_map[gear_position].append(int(number))
                        else:
                            gear_attachment_map[gear_position] = [int(number)]
                    number = ''
                    possible_special_character_indices = []

    for gear_position, attached_part_number_list in gear_attachment_map.items():
        if len(attached_part_number_list) == 2:
            valid_numbers.append(attached_part_number_list[0] * attached_part_number_list[1])

    print(sum(valid_numbers))                  




