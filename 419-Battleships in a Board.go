package leetcode

func countBattleships(board [][]byte) int {
	var m, n int
	if m = len(board); m == 0 { // check input
		return 0
	}
	if n = len(board[0]); n == 0 { // check input
		return 0
	}

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				count++
				if i > 0 && board[i-1][j] == 'X' { // should not count
					count--
				}
				if j > 0 && board[i][j-1] == 'X' { // should not count
					count--
				}
			}
		}
	}

	return count
}
