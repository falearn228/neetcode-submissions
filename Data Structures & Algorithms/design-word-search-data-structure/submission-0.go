
type Node struct {
	childs [26]*Node
	isEnd  bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	var curr = this.root
	for i := range word {
		ch := word[i]
		idx := ch-'a'
		if curr.childs[idx] == nil {
			curr.childs[idx] = &Node{}
		}
		curr = curr.childs[idx]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var recur func(j int, n *Node) bool
	recur = func(j int, n *Node) bool {
		cur := n
		for i := j; i < len(word); i++ {
			c := word[i]
			if c == '.' {
				for _, child := range cur.childs {
					if child != nil && recur(i+1, child) {
						return true
					} 
				}
				return false
			} else {
				idx := c-'a'
				if cur.childs[idx] == nil {
					return false
				}
				cur = cur.childs[idx]
			}
		}
		return cur.isEnd
	}
	
	return recur(0, this.root)
}
