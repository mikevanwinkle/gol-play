import sys

def main():
    """ Entry Function """

    alive = []
    board = setup_board(8)

    lines = sys.stdin.readlines()
    header = lines.pop(0)
    if header.strip() != "#Life 1.06":
        raise Exception("Invalid format")
    for line in lines:
        alive.append([int(i) for i in line.strip().split(" ")])

    for point in alive:
        print(point)
        board[point[0]][point[1]] = 1
    print(board)
    for i in range(0, 10):
        board, alive = update_board(board)
    print(alive)

def update_board(board):
    new_alive = []
    for i in range(0, len(board)):
        for j in range(0, len(board[i])):
            count = count_neighbors(board, i, j)
            if board[i][j] == 0:
                # if this was a dead cell
                if count == 3:
                    new_alive.append([i, j])
                    board[i][j] = 1
            elif board[i][j] == 1:
                if count <2 or count>3:
                    board[i][j] = 0
                else:
                    new_alive.append([i, j])

    return board, new_alive

def count_neighbors(board, x, y):
    neighbors = [
        [-1, -1], [-1, 0], [-1, 1],
        [0, -1], [0, 1],
        [1, -1], [1, 0], [1, 1]
    ]
    count = 0
    for nb in neighbors:
        try:
            # print(x + nb[0], y + nb[1])
            count = count + board[x + nb[0]][y + nb[1]]
        except Exception as e:
            # print(e)
            pass


    return count

def setup_board(size):
    board = []
    for i in range(0, size):
        row = []
        for j in range(0, size): row.append(0)
        board.append(row)

    return board
if __name__ == '__main__':
    main()
