/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	vtd := make(map[*Node]*Node)
    var recur func(curr *Node) *Node
	recur = func(curr *Node) *Node {
		if curr == nil {
			return nil
		}
		if clone, ok := vtd[curr]; ok {
			return clone
		}

		clone := &Node{
			Val: curr.Val,
			Neighbors: make([]*Node, len(curr.Neighbors)),
		}

		vtd[curr] = clone
		for i, nei := range curr.Neighbors {
			clone.Neighbors[i] = recur(nei)
		}
		return clone
	}
	chmo := recur(node)  
	return chmo
}
