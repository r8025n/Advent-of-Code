import copy

input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

map_of_maps = {1: {}, 2: {}, 3:{}, 4:{}, 5:{}, 6:{}, 7:{}}
result = 999999999999

def get_seed_list():
    seed_range_list = []
    seed_line = input_lines[0]
    seed_line = seed_line.split(":")[1].strip()
    seed_pairs = [int(x) for x in seed_line.split()]

    for i in range(0, len(seed_pairs), 2):
        start = seed_pairs[i]
        rrange = seed_pairs[i + 1]
        seed_range_list.append([start, start + rrange - 1])
    return seed_range_list

def generate_map_of_maps():
    map_number = 0

    for i in range(1, len(input_lines)):
        line = input_lines[i]

        if line.isspace():
            map_number += 1
        elif line[0].isnumeric() == False:
            continue
        else:
            destination, source, number_range = [int(x) for x in line.split()]
            source_range = (source, source + number_range)
            difference_between_destination_and_source = destination - source
            map_of_maps[map_number][source_range] = difference_between_destination_and_source
    return

if __name__=="__main__":
    ranges_for_next_level = []
    seed_range_list = get_seed_list()
    generate_map_of_maps()

    for i in range(1, 8):
        ranges_for_next_level = []
        current_map = map_of_maps[i]
        for seed_range in seed_range_list:
            found = False
            for rrange, difference in current_map.items():
                if seed_range[1] > rrange[0] and seed_range[1] < rrange[1]:
                    found = True
                    new_range = [max(seed_range[0], rrange[0]) + difference, seed_range[1] + difference]
                    ranges_for_next_level.append(new_range)
                    if seed_range[0] < rrange[0]:
                        extra_range =  [seed_range[0], rrange[0] - 1]
                        seed_range_list.append(extra_range)
                    break
            if found == False:
                ranges_for_next_level.append(seed_range)
        seed_range_list = copy.deepcopy(ranges_for_next_level)

    
    for seed_range in seed_range_list:
        result = min(result, seed_range[0])

    print(result)

