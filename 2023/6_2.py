input_file = open('input.txt', 'r')
input_lines = input_file.readlines()


if __name__=="__main__":
    time_list = input_lines[0].split(":")[1].strip().split()
    distance_list = input_lines[1].split(":")[1].strip().split()

    time = int(''.join(time_list))
    distance = int(''.join(distance_list))
    print(time, distance)

    count = 0

    for i in range (1, time):
        if i * (time - i) > distance:
            count+=1

    print(count)