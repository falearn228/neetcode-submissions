type Pair [3]int

func islandsAndTreasure(grid [][]int) {
    queue := make([]Pair, 0, len(grid)*len(grid[0]))
    
    for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 0 {
				queue = append(queue, Pair{r, c, 0})
			}
		}
	}

	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

    shift := 0
    for len(queue)-shift > 0 {
        cur := queue[shift]
        shift++

        r, c, step := cur[0], cur[1], cur[2]

        for _, d := range dirs {
            nr, nc := d[0]+r, d[1]+c

            if !check(nr, nc, grid) {
                continue
            }

            if grid[nr][nc] != math.MaxInt32 {
                continue
            }
            grid[nr][nc] = step+1
            queue = append(queue, Pair{nr, nc, step+1})
        }
    }
}

func check(r, c int, board [][]int) bool {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}
	return true
}
