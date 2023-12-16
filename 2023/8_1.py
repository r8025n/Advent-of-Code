input_file = open("input.txt", "r")
input_lines = input_file.readlines()

mmap = {}

if __name__ == "__main__":
    is_command = True
    commands = ""
    count = 0
    command_count = 0

    for line in input_lines:
        
        if line.isspace() == True:
            is_command = False
            continue
        if is_command == True:
            commands += line.strip()
        else:
            line_parts = line.split(" = ")
            source = line_parts[0]
            destination_string = line_parts[1][1:len(line_parts[1].strip()) - 1]
            destination = destination_string.split(", ")
            mmap[source] = destination

    source = "AAA"

    while source != "ZZZ":
        if command_count == len(commands):
            command_count = 0
        if commands[command_count] == "L":
            source = mmap[source][0]
        else:
            source = mmap[source][1]
        
        count += 1
        command_count += 1
    
    print(count)

        