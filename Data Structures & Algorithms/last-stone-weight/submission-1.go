type KthLargest struct {
	arr Heap
}

type Heap []int

func (k Heap) Len() int           { return len(k) }
func (k Heap) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k Heap) Less(i, j int) bool { return k[i] > k[j] }
func (k *Heap) Push(x any)        { *k = append(*k, x.(int)) }
func (k *Heap) Pop() any {
	x := (*k)[len(*k)-1]
	*k = (*k)[:len(*k)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	if len(stones) < 2 {
		return stones[0]
	}
	k := &KthLargest{}
	heap.Init(&k.arr)

	for i := range stones {
		heap.Push(&k.arr, stones[i])
	}
	for {
		s1 := heap.Pop(&k.arr).(int)
		s2 := heap.Pop(&k.arr).(int)

		mass := abs(s1-s2)
		if mass != 0 {
			heap.Push(&k.arr, mass)
		}
		if k.arr.Len() == 1 {
			return k.arr[0] 
		} else if k.arr.Len() == 0 {
			return 0
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
