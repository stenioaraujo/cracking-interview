def main():
    matrix = [
        [1,1,1,1],
        [2,2,2,2],
        [3,3,3,3],
        [4,4,4,4]
    ]
    print_matrix(matrix)

    matrix = [
        [1,1,1],
        [2,2,2],
        [3,3,3]
    ]
    print_matrix(matrix)

    matrix = [
        [1,1],
        [2,2]
    ]
    print_matrix(matrix)

    matrix = [
        [1]
    ]
    print_matrix(matrix)

def rotate(matrix):
    middle = int(len(matrix) / 2)
    for i in range(middle):
        rotate_matrix(matrix, i, len(matrix) - i - 1)

def rotate_matrix(matrix, start, end):
    for i in range(end - start):
        # top with left exchange
        matrix[start][start+i], matrix[end-i][start] = matrix[end-i][start], matrix[start][start+i]

        # right with top exchange
        matrix[start+i][end], matrix[start][start+i] = matrix[start][start+i], matrix[start+i][end]

        # botton with right exchange
        matrix[end][end-i], matrix[start+i][end] = matrix[start+i][end], matrix[end][end-i]

def print_matrix(matrix):
    for l in matrix:
        print(l)
    print("--------------------")
    rotate(matrix)
    for l in matrix:
        print(l)
    print("===================+")

main()