// UVa 10196 - Check The Check

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int }

func find(piece byte, board [8][8]byte) point {
	for i, vi := range board {
		for j, vj := range vi {
			if vj == piece {
				return point{i, j}
			}
		}
	}
	return point{}
}

func inBoard(x, y int) bool { return x >= 0 && x <= 7 && y >= 0 && y <= 7 }

func checkForPieces(king point, board [8][8]byte, pieces []byte, directions [4][2]int) bool {
	for _, direction := range directions {
		x, y := king.x, king.y
		for {
			x += direction[0]
			y += direction[1]
			if !inBoard(x, y) {
				break
			}
			if board[x][y] != '.' {
				for _, piece := range pieces {
					if board[x][y] == piece {
						return true
					}
				}
				break
			}
		}
	}
	return false
}

func checkBishopAndQueen(king point, board [8][8]byte, pieces []byte) bool {
	directions := [4][2]int{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}}
	return checkForPieces(king, board, pieces, directions)
}

func checkRookAndQueen(king point, board [8][8]byte, pieces []byte) bool {
	directions := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	return checkForPieces(king, board, pieces, directions)
}

func checkForPiece(king point, board [8][8]byte, piece byte, directions [][2]int) bool {
	for _, direction := range directions {
		x := king.x + direction[0]
		y := king.y + direction[1]
		if inBoard(x, y) && board[x][y] == piece {
			return true
		}
	}
	return false
}

func checkPawn(king point, board [8][8]byte, direction int, piece byte) bool {
	directions := [][2]int{{-1, direction}, {1, direction}}
	return checkForPiece(king, board, piece, directions)
}

func checkKnight(king point, board [8][8]byte, piece byte) bool {
	directions := [][2]int{{1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}}
	return checkForPiece(king, board, piece, directions)
}

func checkWhite(board [8][8]byte) bool {
	king := find('K', board)
	return checkRookAndQueen(king, board, []byte{'r', 'q'}) ||
		checkBishopAndQueen(king, board, []byte{'b', 'q'}) ||
		checkPawn(king, board, -1, 'p') ||
		checkKnight(king, board, 'n')
}

func checkBlack(board [8][8]byte) bool {
	king := find('k', board)
	return checkRookAndQueen(king, board, []byte{'R', 'Q'}) ||
		checkBishopAndQueen(king, board, []byte{'B', 'Q'}) ||
		checkPawn(king, board, 1, 'P') ||
		checkKnight(king, board, 'N')
}

func check(out *os.File, kase int, board [8][8]byte) {
	fmt.Fprintf(out, "Game #%d: ", kase)
	switch {
	case checkWhite(board):
		fmt.Fprintln(out, "white king is in check.")
	case checkBlack(board):
		fmt.Fprintln(out, "black king is in check.")
	default:
		fmt.Fprintln(out, "no king is in check.")
	}
}

func main() {
	in, _ := os.Open("10196.in")
	defer in.Close()
	out, _ := os.Create("10196.out")
	defer out.Close()

	var line string
	var board [8][8]byte
	for kase := 1; ; kase++ {
		done := true
		for i := range board {
			fmt.Fscanf(in, "%s", &line)
			for j := range board[i] {
				board[i][j] = line[j]
				if line[j] != '.' {
					done = false
				}
			}
		}
		if done {
			break
		}
		fmt.Fscanln(in)
		check(out, kase, board)
	}
}
