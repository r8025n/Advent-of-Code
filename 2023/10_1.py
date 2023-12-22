from collections import deque

input_file = open("input.txt", "r")
input_lines = input_file.readlines()

memory = []
adjacency_map = {
    (0, 1): ["-SFL", "-SJ7"],
    (0, -1): ["-SJ7", "-SFL"],
    (1, 0): ["|S7F", "|SJL"],
    (-1, 0): ["|SJL", "|S7F"]
}

xxx = [-1, 0, 0, 1]
yyy = [0, -1, 1, 0]

def valid_connection(connection, adjacency):
    if connection[0] in adjacency_map[adjacency][0] and connection[1] in adjacency_map[adjacency][1]:
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
        
                if valid_connection(connection, (xxx[i], yyy[i])) and memory[xx][yy] == -1:
                    qqueue.append((xx, yy))
                    memory[xx][yy] = memory[x][y] + 1                
                    max_distance = max(max_distance, memory[xx][yy])

    # with open('output.txt', 'w') as file:
    #     for mem in memory:
    #         file.write(",".join([str(x) for x in mem]) + "\n")
    print(max_distance)


