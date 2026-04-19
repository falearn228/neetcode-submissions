type KthLargest struct {
	arr Heap
}

type Heap [][]int

func (k Heap) Len() int { return len(k) }
func (k Heap) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k Heap) Less(i, j int) bool {
	di := dist(float64(k[i][0]), float64(k[i][1]))
	dj := dist(float64(k[j][0]), float64(k[j][1]))
	if di > dj {
		return true
	}
	return false
}
func (k *Heap) Push(x any) { *k = append(*k, x.([]int)) }
func (k *Heap) Pop() any {
	x := (*k)[len(*k)-1]
	*k = (*k)[:len(*k)-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &KthLargest{}
	heap.Init(&h.arr)

	for i := range points {

		heap.Push(&h.arr, points[i])
		if h.arr.Len() > k {
			heap.Pop(&h.arr)
		}
	}

	return h.arr
}

func dist(x, y float64) float64 {
	d := x*x + y*y
	return math.Sqrt(d)
}

