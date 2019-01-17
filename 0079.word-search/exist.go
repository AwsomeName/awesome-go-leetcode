package problem0079

func exist(board [][]byte, word string) bool {
	// find := false
	mark := make([][]bool, len(board))
	for i := range mark {
		mark[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				mark[i][j] = true
				if search(board, word[1:], mark, i, j) {
					return true
				}
				mark[i][j] = false
			}
		}

	}
	return false
}

func search(board [][]byte, word string, mark [][]bool, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i+1 < len(board) && board[i+1][j] == word[0] && !mark[i+1][j] {
		mark[i+1][j] = true
		if search(board, word[1:], mark, i+1, j) {
			return true
		}
		mark[i+1][j] = false
	}
	if i-1 >= 0 && board[i-1][j] == word[0] && !mark[i-1][j] {
		mark[i-1][j] = true
		if search(board, word[1:], mark, i-1, j) {
			return true
		}
		mark[i-1][j] = false
	}
	if j+1 < len(board[0]) && board[i][j+1] == word[0] && !mark[i][j+1] {
		mark[i][j+1] = true
		if search(board, word[1:], mark, i, j+1) {
			return true
		}
		mark[i][j+1] = false
	}
	if j-1 >= 0 && board[i][j-1] == word[0] && !mark[i][j-1] {
		mark[i][j-1] = true
		if search(board, word[1:], mark, i, j-1) {
			return true
		}
		mark[i][j-1] = false
	}
	return false
}
