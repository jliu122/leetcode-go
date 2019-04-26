package heap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Val int
	Freq int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int  {
	return len(h)
}

func (h ItemHeap) Less(i, j int) bool {
	return h[i].Freq > h[j].Freq
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ItemHeap) Push(x interface{}) {
	node := x.(*Item)
	*h = append(*h, node)
}

func (h *ItemHeap) Pop() interface{}  {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[: n - 1]
	return x
}

// 347
func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)

	h := &ItemHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, &Item{k, v})
	}

	result := make([]int, k)Å
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(*Item).Val
	}
	return result
}
