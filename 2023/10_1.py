from collections import deque

input_file = open("input.txt", "r")
input_lines = input_file.readlines()

memory = []
adjacency_map = {
    (0, 1): ["--", "F-", "L-", "-7", "-J", "LJ", "F7", "FJ", "L7"],
    (0, -1): ["--", "-F", "-L", "7-", "J-","JL", "7F", "JF", "7L"],
    (1, 0): ["||", "F|", "|J", "|L", "7|", "FJ", "FL", "7L", "7J"],
    (-1, 0): ["||", "|F", "J|", "L|", "|7", "JF", "LF", "L7", "J7"]
}

xxx = [-1, 0, 0, 1]
yyy = [0, -1, 1, 0]

def valid_connection(con):
    valid_characters = ["F", "J", "7", "L", "-", "|"]
    if con[0] == "S" and con[1] in valid_characters:
        return True
    if con[1] == "S" and con[0] in valid_characters:
        return True

    return False 

def find_starting_point(lines):
    for i in range(0, len(lines)):
        for j in range(0, len(lines[i])):
            if lines[i][j] == "S":
                return (i, j)

if __name__ == "__main__":
    for line in input_lines:
        new_list = []
        for char in line:
            new_list.append(-1)
        memory.append(new_list)

    max_distance = 0
    qqueue = deque()
    starting_point = find_starting_point(input_lines)
    memory[starting_point[0]][starting_point[1]] = 0
    qqueue.append(starting_point)
    while len(qqueue) != 0:
        x, y = qqueue.popleft()
        for i in range(0, 4):
            xx = x + xxx[i]
            yy = y + yyy[i]
            if xx >= 0 and xx < len(input_lines) and yy >= 0 and yy < len(input_lines[-1]):
                connection = input_lines[x][y] + input_lines[xx][yy]
                if (connection in adjacency_map[(xxx[i], yyy[i])] or valid_connection(connection) == True) and memory[xx][yy] == -1:
                    # print("////////")
                    qqueue.append((xx, yy))
                    memory[xx][yy] = memory[x][y] + 1                
                    max_distance = max(max_distance, memory[xx][yy])

    with open('output.txt', 'w') as file:
        for mem in memory:
            file.write(",".join([str(x) for x in mem]) + "\n")
    print(max_distance)


