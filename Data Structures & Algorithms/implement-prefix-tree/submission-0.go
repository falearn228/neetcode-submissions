type Node struct {
	childs [26]*Node
	isEnd  bool
}

type PrefixTree struct {
	tree *Node
}

func Constructor() PrefixTree {
	return PrefixTree{
		tree: &Node{},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.tree
	for i := range word {
		ch := word[i]
		idx := ch - 'a'
		if curr.childs[idx] == nil {
			curr.childs[idx] = &Node{}
		}
		curr = curr.childs[idx]
	}
	curr.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.tree
	for i := range word {
		ch := word[i]
		idx := ch - 'a'
		if curr.childs[idx] == nil {
			return false
		}
		curr = curr.childs[idx]
	}
	return curr.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.tree
	for i := range prefix {
		ch := prefix[i]
		idx := ch - 'a'
		if curr.childs[idx] == nil {
			return false
		}
		curr = curr.childs[idx]
	}
	return true
}
