
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var ans int
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if !check(i, j, grid) {
			return false
		}

		grid[i][j] = '#'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)

		return true
	}
	
	for i := range rows {
		for j := range cols {
			if dfs(i, j) {
				ans++
			}
		}
	}
	return ans
}

func check(r, c int, board [][]byte) bool {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || 
	board[r][c] == '0' ||
	board[r][c] == '#' {
		return false
	} 
	return true
}
