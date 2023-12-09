input_file = open('input.txt', 'r')
input_lines = input_file.readlines()

final_outputs = []

if __name__=="__main__":
    time_list_in_string = input_lines[0].split(":")[1].strip().split()
    distance_list_in_string = input_lines[1].split(":")[1].strip().split()

    time_list = [int(time) for time in input_lines[0].split(":")[1].strip().split()]
    distance_list = [int(distance) for distance in input_lines[1].split(":")[1].strip().split()]

    for i in range(0, len(time_list)):
        count = 0
        time = time_list[i]
        distance = distance_list[i]

        for i in range(1, distance - 1):
            if i * (time- i) > distance:
                count+=1

        final_outputs.append(count)

    output = 1

    for val in final_outputs:
        output = output * val
    
    print(output)