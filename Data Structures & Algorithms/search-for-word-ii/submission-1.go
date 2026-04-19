

type Node struct {
	childs [26]*Node
	idx    int
}

func Constructor() *Node {
	return &Node{
		idx: -1,
	}
}

func (this *Node) addWord(word string, i int) {
	cur := this
	for _, ch := range word {
		idx := ch - 'a'
		if cur.childs[idx] == nil {
			cur.childs[idx] = Constructor() 
		}
		cur = cur.childs[idx]
	}
	cur.idx = i
}

func findWords(board [][]byte, words []string) []string {
	root := Constructor()
	for i, word := range words {
		root.addWord(word, i)
	}

	rows, cols := len(board), len(board[0])
	var res []string

	getIdx := func(c byte) int { return int(c - 'a') }

	var dfs func(r, c int, node *Node)
	dfs = func(r, c int, node *Node) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] == '#' ||
		 node.childs[getIdx(board[r][c])] == nil {
			return
		}

		tmp := board[r][c]
		board[r][c] = '#'
		node = node.childs[getIdx(tmp)]
		if node.idx != -1 {
			res = append(res, words[node.idx])
			node.idx = -1
		}

		dfs(r+1, c, node)
		dfs(r-1, c, node)
		dfs(r, c+1, node)
		dfs(r, c-1, node)

		board[r][c] = tmp
	}

	for i := range rows {
		for j := range cols {
			dfs(i, j, root)
		}
	}
	return res
}
