import math

input_file = open("input.txt", "r")
input_lines = input_file.readlines()

mmap = {}

def concurrent_destination_list_check(destination_list):
    # print(destination_list)
    for dest in destination_list:
        if dest[-1] != "Z":
            return False

    return True

def lcm(a, b):
    return abs(a * b) // math.gcd(a, b)

def lcm_of_list(numbers):
    result = 1
    for num in numbers:
        result = lcm(result, num)
    return result

if __name__ == "__main__":
    is_command = True
    commands = ""
    count = 0
    command_count = 0
    source_list = []
    count_list = []

    for line in input_lines:
        
        if line.isspace() == True:
            is_command = False
            continue
        if is_command == True:
            commands += line.strip()
        else:
            line_parts = line.split(" = ")
            source = line_parts[0]
            if source[-1] == "A":
                source_list.append(source)
            destination_string = line_parts[1][1:len(line_parts[1].strip()) - 1]
            destination = destination_string.split(", ")
            mmap[source] = destination
    
    for source in source_list:
        count = 0
        command_count
        while source[-1] != "Z":
            if command_count == len(commands):
                    command_count = 0
            if commands[command_count] == "L":
                source = mmap[source][0]
            else:
                source = mmap[source][1]
            
            count += 1
            command_count += 1

        count_list.append(count)
        
    


    print(lcm_of_list(count_list))
        