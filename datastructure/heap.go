package datastructure

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *IntHeap) Swap(i, j int)      { *h[i], *h[j] = *h[j], *h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An String is a min-heap of strings.
type StringHeap []string

func (h StringHeap) Len() int {
	return len(h)
}

func (h StringHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h StringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StringHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[:n - 1]
	return x
}