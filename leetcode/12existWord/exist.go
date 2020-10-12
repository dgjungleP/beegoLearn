package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if foundNext(i, j, 0, board, word) {
					return foundNext(i, j, 0, board, word)
				}
			}
		}
	}
	return false
}

func foundNext(cow int, cloumn int, k int, board [][]byte, word string) bool {
	if k >= len(word) {
		return true
	}
	if cow < 0 || cow >= len(board) || cloumn < 0 || cloumn >= len(board[cow]) {
		return false
	}
	if board[cow][cloumn] == word[k] {
		board[cow][cloumn] = ' '
		k++
		result := foundNext(cow+1, cloumn, k, board, word) ||
			foundNext(cow-1, cloumn, k, board, word) ||
			foundNext(cow, cloumn+1, k, board, word) ||
			foundNext(cow, cloumn-1, k, board, word)
		board[cow][cloumn] = word[k-1]
		return result
	}
	return false
}
func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// board := [][]byte{{'a', 'a'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
