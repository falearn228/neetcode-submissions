/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
import (
	"slices"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// Input: intervals = [(0,30),(5,10),(15,20)]
func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(x, y Interval) int {
		return x.start - y.start
	})

	var ans int

	h := IntHeap{}
	heap.Init(&h)

	for _, it := range intervals {
		for h.Len() > 0 && h[0] <= it.start {
			heap.Pop(&h)
		}

		heap.Push(&h, it.end)
		ans = max(ans, h.Len())
	}

	return ans
}
