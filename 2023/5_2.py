input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

map_of_maps = {1: {}, 2: {}, 3:{}, 4:{}, 5:{}, 6:{}, 7:{}}
map_number = 0
location_list = []

def recursive_approach(level, seed_list):
    for val in seed_list:
        list_for_next_level = []
        in_range = False
        for rrange, difference in map_of_maps[level].items():
            if val >= rrange[1] and val <= rrange[2]:
                list_for_next_level.append(val + difference)
                in_range = True
                
        if in_range == False:
            list_for_next_level.append(val)
            in_range = False
        if (level + 1) <= 7:
            recursive_approach(level + 1, list_for_next_level)
        else:
            location_list.append(min(list_for_next_level))

    return

def get_seed_list(seed_pairs):
    seeds = []
    for i in range(0, len(seed_pairs), 2):
        start = seed_pairs[i]
        rrange = seed_pairs[i + 1]
        for i in range(0, rrange + 1):
            seeds.append(start + i)
    # print("/////////////////////")        
    # print(len(seeds))
    # print(len(set(seeds)))
    print("/////////////////////")
    return seeds


if __name__=="__main__":
    seed_line = input_lines[0]
    seed_line = seed_line.split(":")[1].strip()

    seed_pairs = [int(x) for x in seed_line.split()]

    seeds = get_seed_list(seed_pairs)

    for i in range(1, len(input_lines)):
        line = input_lines[i]

        if line.isspace():
            map_number += 1
        elif line[0].isnumeric() == False:
            continue
        else:
            destination, source, number_range = [int(x) for x in line.split()]
            source_range = (i, source, source + number_range) #i is added to make the key unique
            difference_between_destination_and_source = destination - source
            map_of_maps[map_number][source_range] = difference_between_destination_and_source

    recursive_approach(1, seeds)

    print(min(location_list))