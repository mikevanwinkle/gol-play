import sys

def main():
    """ Entry Function """
    lines = sys.stdin.readlines()
    header = lines.pop(0)
    raw_points = []
    # check header
    if header.strip() != "#Life 1.06":
        raise Exception("Invalid format")
    for line in lines:
        raw_points.append([int(i) for i in line.strip().split(" ")])

    size, nums = calc_board_size(raw_points)
    board = setup_board(size)
    offset = sorted(nums)[0]
    alive = []
    for point in raw_points:
        alive.append([point[0]-offset, point[1]-offset])

    for point in alive:
        board[point[0]][point[1]] = 1

    for i in range(0, 10):
        board, alive = update_board(board)

    print("#Life 1.06")
    for point in alive:
        print(f"{point[0] + offset} {point[1] + offset}")

def calc_board_size(points):
    nums = []
    for pair in points:
        for n in pair:
            nums.append(n)
    mn, mx = min(nums), max(nums)
    root = (mn - mx) + ((mn - mx) % 2)
    return abs(root * root * root), nums

def update_board(board):
    new_alive = []
    new_board = board
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
