

type Node struct {
	childs [26]*Node
	idx    int
	refs   int
}

func Constructor() *Node {
	return &Node{
		idx: -1,
	}
}

func (this *Node) addWord(word string, i int) {
	cur := this
	cur.refs++
	for _, ch := range word {
		idx := ch - 'a'
		if cur.childs[idx] == nil {
			cur.childs[idx] = Constructor() 
		}
		cur = cur.childs[idx]
		cur.refs++
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
		prev := node
		node = node.childs[getIdx(tmp)]
		if node.idx != -1 {
			res = append(res, words[node.idx])
			node.idx = -1
			node.refs--
			if node.refs == 0 {
				prev.childs[getIdx(tmp)] = nil
				board[r][c] = tmp
				return
			}
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
