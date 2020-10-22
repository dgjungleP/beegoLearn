package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]int, 9)
	box := make([]map[byte]int, 9)
	cloumn := make([]map[byte]int, 9)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if _, ok := row[i][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					if row[i] == nil {
						row[i] = make(map[byte]int)
					}
					row[i][board[i][j]] = 1
				}

			}
			if _, ok := cloumn[j][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					if cloumn[j] == nil {
						cloumn[j] = make(map[byte]int)
					}
					cloumn[j][board[i][j]] = 1
				}
			}
			if _, ok := box[(i/3)*3+j/3][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					if box[(i/3)*3+j/3] == nil {
						box[(i/3)*3+j/3] = make(map[byte]int)
					}
					box[(i/3)*3+j/3][board[i][j]] = 1
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Print(isValidSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'5', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
}
