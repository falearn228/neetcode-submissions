func isValidSudoku(board [][]byte) bool {
  rows := make([]map[byte]bool, 9)
  cols := make([]map[byte]bool, 9)
  squares := make([]map[byte]bool, 9)

  for i := range rows {
    rows[i] = make(map[byte]bool)
    cols[i] = make(map[byte]bool)
    squares[i] = make(map[byte]bool)
  }

  for i := range rows {
    for j := range cols {
      if board[i][j] == '.' {
        continue
      }

      ch := board[i][j]
      square := (i/3)*3 + j/3 
  
      if rows[i][ch] || cols[j][ch] || squares[square][ch] {
        return false
      } 

      rows[i][ch] = true
      cols[j][ch] = true
      squares[square][ch] = true
    }
  }

  return true
}
