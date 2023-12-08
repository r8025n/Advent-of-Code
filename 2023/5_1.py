input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

INT_MAX = 99999999999
source_vs_difference_map = {}

if __name__=="__main__":
    seed_line = input_lines[0]
    seed_line = seed_line.split(":")[1].strip()

    seeds = [int(x) for x in seed_line.split()]

    for i in range(1, len(input_lines)):
        line = input_lines[i]

        if line.isspace():
            print("seeds_before",seeds)
            # print(source_vs_difference_map)
            min_value = INT_MAX

            for i in range(0, len(seeds)):
                for rrange, difference in source_vs_difference_map.items():
                    if seeds[i] >= rrange[1] and seeds[i] <= rrange[2]:
                        min_value = min(seeds[i] + difference, min_value)

                if min_value != INT_MAX:
                    seeds[i] = min_value
                    min_value = INT_MAX
            # print("seeds_after",seeds)
            source_vs_difference_map = {}
        elif line[0].isnumeric() == False:
            continue
        else:
            destination, source, number_range = [int(x) for x in line.split()]
            source_range = (i, source, source + number_range) #i is added to make the key unique
            difference_between_destination_and_source = destination - source
            source_vs_difference_map[source_range] = difference_between_destination_and_source
    # print(seeds)
    print(min(seeds))