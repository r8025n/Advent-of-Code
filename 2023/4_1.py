input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

elfs_points = 0

def process_card(line):
    splitted_line = line.split("|")
    splitted_line[0] = splitted_line[0].strip()
    splitted_line[1] = splitted_line[1].strip()

    winning_number_string = splitted_line[0].split(":")[1].strip()
    elfs_number_string = splitted_line[1].strip()

    winning_numbers = [int(x) for x in winning_number_string.split()]
    elfs_numbers = [int(x) for x in  elfs_number_string.split()]

    return (winning_numbers, elfs_numbers)

def binary_search_number_in_list(number_list, target):
    low, high = 0, len(number_list) - 1

    while low <= high:
        mid = (low + high) // 2

        if number_list[mid] == target:
            return True
        elif target < number_list[mid]:
            high = mid - 1
        else:
            low = mid + 1
    
    return False

def find_elfs_number_in_winning_number(winning_numbers, elfs_numbers):
    count = 0
    winning_numbers.sort()

    for number in elfs_numbers:
        if number <= winning_numbers[-1] and binary_search_number_in_list(winning_numbers, number):
            count += 1

    return count        


if __name__=="__main__": 

    for line in input_lines:
        winning_numbers, elfs_numbers = process_card(line)

        elfs_number_count_in_winning_number = find_elfs_number_in_winning_number(winning_numbers, elfs_numbers)
        if elfs_number_count_in_winning_number <= 1:
            elfs_points += elfs_number_count_in_winning_number
        else:
            elfs_points += 2 ** (elfs_number_count_in_winning_number - 1)

    print(elfs_points)

