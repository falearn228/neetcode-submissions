type chlen struct {
	cost int
	ind  int
}

func networkDelayTime(times [][]int, n int, k int) int {
	var ans, seen int
	h := MinHeap{}
	heap.Init(&h)

	mapp := make(map[int][]chlen)
	for i := range times {
		mapp[times[i][0]] = append(mapp[times[i][0]], chlen{ind: times[i][1], cost: times[i][2]})
	}
	heap.Push(&h, chlen{ind: k, cost: 0})
	visited := make([]bool, n+1)

	for h.Len() > 0 {
		node := heap.Pop(&h).(chlen)
		if visited[node.ind] {
			continue
		}
		visited[node.ind] = true
		seen++
		ans = node.cost

		for _, neigs := range mapp[node.ind] {
			if !visited[neigs.ind] {
				heap.Push(&h, chlen{
					ind: neigs.ind,
					cost: neigs.cost + node.cost,
				})
			}
		}
	}

	if seen != n {
		return -1
	}

	return ans
}

type MinHeap []chlen

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(chlen))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	old = old[:n]
	*h = old
	return x
}
