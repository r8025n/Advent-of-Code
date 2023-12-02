max_red = 12
max_green = 13
max_blue = 14

input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

necessary_powers = []

for line in input_lines:
    red_count = -1
    green_count = -1
    blue_count = -1

    splitted_line = line.split(': ')
    combination_set = splitted_line[1].split('; ')
    game_id = int(splitted_line[0].split(' ')[1])

    for combination in combination_set:
        combination = combination.strip()
        seperate_values_in_each_combination = combination.split(', ')
        for value in seperate_values_in_each_combination:
            value = value.strip()
            count, color = value.split(' ')
            count = int(count)

            if color == 'red':
                red_count = max(count, red_count)
            elif color == 'green':
                green_count = max(count, green_count)
            else:
                blue_count = max(count, blue_count)

    power = red_count * green_count * blue_count
    necessary_powers.append(power)


print(sum(necessary_powers))
