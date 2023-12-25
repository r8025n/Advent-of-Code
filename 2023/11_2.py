input_file = open("input.txt", "r")
input_lines = input_file.readlines()

galaxy_to_position_map = []
sum_of_shortest_path_length = 0

def get_empty_rows_and_columns(lines):
    empty_rows = [i for i, row in enumerate(lines) if '#' not in row]
    empty_columns = [j for j, column in enumerate(zip(*lines)) if '#' not in column]

    return sorted(empty_rows), sorted(empty_columns)

if __name__ == "__main__":
    for i in range(0, len(input_lines)):
        input_lines[i] = input_lines[i].strip()

    empty_rows, empty_cols = get_empty_rows_and_columns(input_lines)

    for i in range(0, len(input_lines)):
        for j in range(0, len(input_lines[0])):
            if input_lines[i][j] == "#":
                galaxy_to_position_map.append([i, j])

    for i in range(0, len(galaxy_to_position_map)):
        source_x, source_y = galaxy_to_position_map[i]
        for j in range(i + 1, len(galaxy_to_position_map)):
            dest_x, dest_y = galaxy_to_position_map[j]
            expansion_x = len([x for x in empty_rows if x <= max(dest_x, source_x) and x >= min(dest_x, source_x)]) * (1000000 - 1)
            expansion_y = len([y for y in empty_cols if y <= max(dest_y, source_y) and y >= min(dest_y, source_y)]) * (1000000 - 1)
            sum_of_shortest_path_length += (abs(dest_x - source_x) + abs(dest_y - source_y)) + expansion_x + expansion_y

    print(sum_of_shortest_path_length)




   



