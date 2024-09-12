package main

import (
	"bufio"
	"errors"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MAX_INPUT = 100

func main() {
	board, alive := [][]int64{}, [][]int64{}
	alive, err := LoadInput()
	// I've chosed to start by calculting board size based on inputs
	// This will accommodate large numbers that are in the same neighborhood
	// however to accommodate large numbers that are not in the same neighborhood
	// we would need to find some way to carve the space up into multiple boards
	min, max := CalcBoardSize(alive)
	root := (min - max) + ((min - max) % 2)
	size := root * root
	slog.Info("Creating board", "size", size, "root", root)
	board = CreateBoard(size)
	board = SetBoard(board, alive, min)
	for i := 0; i < 10; i++ {
		slog.Info("Iteration ", "i", i)
		board, alive = UpdateBoard(board, alive)
	}

	if err != nil {
		slog.Error(err.Error())
	}

	if err != nil {
		panic(err)
	}
}

func SetBoard(board [][]int64, living [][]int64, offset int64) [][]int64 {
	for _, member := range living {
		board[member[0]][member[1]] = 1
	}

	return board
}

func UpdateBoard(board [][]int64, alive [][]int64) ([][]int64, [][]int64) {
	updateAlive := [][]int64{}
	updatedBoard := board
	for i := int64(0); i < int64(len(board)); i++ {
		for j := int64(0); j < int64(len(board)); j++ {
			print("Count", i, j)
			count := CountNeighbors(board, i, j)
			if board[i][j] == 1 {
				if count < 2 || count > 3 {
					// die
					updatedBoard[i][j] = 0
				} else {
					updateAlive = append(updateAlive, []int64{int64(i), int64(j)})
					updatedBoard[i][j] = board[i][j]
				}
			} else if board[i][j] == 0 && count == 3 {
				// live
				updatedBoard[i][j] = 1
				updateAlive = append(updateAlive, []int64{int64(i), int64(j)})
			}
		}
	}

	return updatedBoard, updateAlive
}

func CalcBoardSize(living [][]int64) (int64, int64) {
	ints := []int{}
	for _, row := range living {
		for _, col := range row {
			ints = append(ints, int(col))
		}
	}
	sort.Ints(ints)

	return int64(ints[0]), int64(ints[len(ints)-1])
}

func CountNeighbors(board [][]int64, i int64, j int64) int64 {
	print("test", board, i, j, "\n")
	neighbors := [][]int64{
		{-1, 1}, {-1, 0}, {-1, -1},
		{0, 1}, {0, 1},
		{1, -1}, {1, 0}, {1, -1},
	}
	count := int64(0)
	for _, nb := range neighbors {
		slog.Info("Neighbor", "nb", nb, "t", int64(len(board)) > i-1)

		if nb[0] == -1 && i > 0 {
			// count the previous row's neighbors
			count = count + board[i+nb[0]][j+nb[1]]
		} else if nb[0] == 1 && int64(len(board)) >= i+1 {
			count = count + board[i+nb[0]][j+nb[1]]
		} else {
			if j+nb[1] < int64(len(board)) {
				count = count + board[i+nb[0]][j+nb[1]]
			}

		}
	}

	return count
}

func LoadInput() ([][]int64, error) {
	reader := bufio.NewReader(os.Stdin)
	first, _ := reader.ReadByte()
	if string(first) != "#" {
		return nil, errors.New("invalid format")
	}

	data, err := reader.ReadBytes('\n')
	if err != nil || string(data) != "Life 1.06\n" {
		return nil, errors.New("Invalid header")
	}
	living := [][]int64{}
	for i := 0; i < MAX_INPUT; i++ {
		slog.Info("reading i")
		xstr, err := reader.ReadString(' ')
		if err != nil {
			slog.Error(err.Error())
			break
		}
		ystr, err := reader.ReadString('\n')
		if err != nil {
			slog.Error(err.Error())
			break
		}
		y, err := strconv.Atoi(strings.Trim(ystr, "\n"))
		if err != nil {
			slog.Error(err.Error())
			break
		}
		x, err := strconv.Atoi(strings.Trim(xstr, " "))
		if err != nil {
			slog.Error(err.Error())
			break
		}
		point := []int64{}
		point = append(point, int64(x))
		point = append(point, int64(y))
		living = append(living, point)
	}
	return living, nil
}

func CreateBoard(size int64) [][]int64 {
	board := make([][]int64, size)
	for i, _ := range board {
		board[i] = make([]int64, size)
	}

	return board
}
