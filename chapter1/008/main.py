def main():
    matrix = [
        [1,1,1,1],
        [2,2,0,2],
        [3,0,3,3],
        [4,4,4,4]
    ]
    print_matrix(matrix)

    matrix = [
        [1,1,1],
        [2,0,2],
        [3,3,3]
    ]
    print_matrix(matrix)

    matrix = [
        [1,1],
        [2,0]
    ]
    print_matrix(matrix)

    matrix = [
        [1]
    ]
    print_matrix(matrix)

    matrix = [
        []
    ]
    print_matrix(matrix)

def zero(matrix):
    if len(matrix) == 0:
        return

    num_lines = len(matrix)
    num_cols = len(matrix[0])

    lines = [False] * num_lines
    columns = [False] * num_cols

    for i in range(num_lines):
        for j in range(num_cols):
            if matrix[i][j] == 0:
                lines[i] = True
                columns[j] = True
    
    for i, should_zero in enumerate(lines):
        if should_zero:
            zeroLine(matrix, i)
    
    for j, should_zero in enumerate(columns):
        if should_zero:
            zeroColumn(matrix, j)

def zeroLine(matrix, i):
    for j in range(len(matrix[0])):
        matrix[i][j] = 0

def zeroColumn(matrix, j):
    for i in range(len(matrix)):
       matrix[i][j] = 0

def print_matrix(matrix):
    for l in matrix:
        print(l)
    print("--------------------")
    zero(matrix)
    for l in matrix:
        print(l)
    print("===================+")

main()