input_file = open("input.txt", "r")
input_lines = input_file.readlines()

def recursive_solution(data_list):
    output = []
    if sum(data_list) == 0:
        return 0
    
    for i in range(1, len(data_list)):
        output.append(data_list[i] - data_list[i - 1])
    
    return data_list[-1] + recursive_solution(output)


if __name__ == "__main__":
    output = []
    for line in input_lines:
        input_list = [int(x) for x in line.split()]
        output.append(recursive_solution(input_list))
    
    print(sum(output))

