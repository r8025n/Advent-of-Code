from collections import deque

input_file = open("input.txt", "r")
input_lines = input_file.readlines()

in_loop_count = 0
memory = []
adjacency_map = {
    (0, 1): ["-SFL", "-SJ7"],
    (0, -1): ["-SJ7", "-SFL"],
    (1, 0): ["|S7F", "|SJL"],
    (-1, 0): ["|SJL", "|S7F"]
}

xxx = [-1, 0, 0, 1]
yyy = [0, -1, 1, 0]

end_point = []

def valid_connection(connection, adjacency):
    if connection[0] in adjacency_map[adjacency][0] and connection[1] in adjacency_map[adjacency][1]:
        return True

    return False 

def find_starting_point(lines):
    for i in range(0, len(lines)):
        for j in range(0, len(lines[i])):
            if lines[i][j] == "S":
                return (i, j)

def check_if_in_loop_by_crossing_count(x, y):
    vertical_crossing_count = 0
    horizantal_crossing_count = 0

    for j in range(y + 1, len(input_lines[x])):
        if input_lines[x][j] in "|J7S":
            vertical_crossing_count += 1
    for i in range(x + 1, len(input_lines)):
        if input_lines[i][y] in "-FLS":
            horizantal_crossing_count += 1

    if vertical_crossing_count % 2 != 0 and horizantal_crossing_count % 2 != 0:
        return True
    
    return False

if __name__ == "__main__":
    for line in input_lines:
        new_list = []
        for char in line:
            new_list.append(-5)
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
        
                if valid_connection(connection, (xxx[i], yyy[i])) and memory[xx][yy] == -5:
                    qqueue.append((xx, yy))
                    memory[xx][yy] = memory[x][y] + 1                
                    max_distance = max(max_distance, memory[xx][yy])
        
        if len(qqueue) == 0:
                    end_point.append(x)
                    end_point.append(y)
    print(end_point)

    in_loop_nodes = []

    # for i in range(0, len(input_lines)):
    #     for j in range(0, len(input_lines[i])):
    #         if memory[i][j] != -1:
    #             memory[i][j] = 0

    m, n = end_point
    qqueue.append((m,n))
    in_loop_nodes.append((m, n))

    while memory[m][n] != 0:
        m, n = qqueue.popleft()
        val = memory[m][n]
        for i in range(0, 4):
            mm = m + xxx[i]
            nn = n + yyy[i]
            if mm >= 0 and mm < len(input_lines) and nn >= 0 and nn < len(input_lines[-1]):
                if memory[mm][nn] == val - 1:
                    qqueue.append((mm, nn))
                    in_loop_nodes.append((mm, nn))

    for node in in_loop_nodes:
        memory[node[0]][node[1]] = 0

    # for i in range(0, len(input_lines)):
    #     loop_cross_count = 1
    #     for j in range(len(input_lines[i]) - 1, -1, -1):
    #         if input_lines[i][j] in "|FJL7S":
    #             memory[i][j] = loop_cross_count
    #             loop_cross_count += 1

    for i in range(0, len(input_lines)):
        for j in range(0, len(input_lines[-1])):
            if memory[i][j] == -5 and check_if_in_loop_by_crossing_count(i, j) == True:
                memory[i][j] = -1
                in_loop_count += 1
        print(in_loop_count)

    with open('output.txt', 'w') as file:
        for mem in memory:
            file.write(",".join([str(x) for x in mem]) + "\n")


    print(max_distance)
    print(in_loop_count)


