type KthLargest struct {
	arr Heap
	k   int
}

type Heap []int

func (k Heap) Len() int           { return len(k) }
func (k Heap) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k Heap) Less(i, j int) bool { return k[i] < k[j] }
func (k *Heap) Push(x any)        { *k = append(*k, x.(int)) }
func (k *Heap) Pop() any {
	x := (*k)[len(*k)-1]
	*k = (*k)[:len(*k)-1]
	return x
}

func Constructor(k int, nums []int) *KthLargest {
	klar := &KthLargest{k: k}
	heap.Init(&klar.arr)

	for _, n := range nums {
		klar.Add(n)
	}

	return klar
}

func (kl *KthLargest) Add(val int) int {
	if kl.arr.Len() < kl.k {
		heap.Push(&kl.arr, val)
	} else if val > kl.arr[0] {
		heap.Pop(&kl.arr)
		heap.Push(&kl.arr, val)
	}
	return kl.arr[0]
}