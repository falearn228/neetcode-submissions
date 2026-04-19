func maxAreaOfIsland(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    var ans, maxAns int
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if !check(r, c, n, m, grid) {
            return 0
        }

        grid[r][c] = 2

        return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
    }

    for r := range n {
        for c := range m {
            if grid[r][c] == 1 {
                ans = dfs(r, c)
            }
            maxAns = max(ans, maxAns)
        }
    }

    return maxAns
}

func check(r, c, n, m int, grid [][]int) bool {
    if r < 0 || c < 0 || r >= n || c >= m || grid[r][c] != 1 {
        return false
    }
    return true
}
