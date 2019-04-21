package graph

import "container/heap"

type StringHeap []string

func (h StringHeap) Len() int  {
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

func (h *StringHeap) Pop() interface{}  {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[: n - 1]
	return x
}

func findItinerary(tickets [][]string) []string {

	m := make(map[string]*StringHeap)
	for _, t := range tickets {
		from, to := t[0], t[1]
		if _, ok := m[from]; ok {
			m[from].Push(to)
		} else {
			h := &StringHeap{}
			heap.Init(h)
			heap.Push(h, to)
			m[from] = h
		}
	}
	var result []string
	dfs("JFK", m, &result)
	return result
}

func dfs(airport string, m map[string]*StringHeap, result *[]string)  {
	if _, ok := m[airport]; ok {
		for m[airport].Len() != 0 {
			next := heap.Pop(m[airport]).(string)
			dfs(next, m, result)
		}
	}
	*result = append([]string{airport}, *result...)
}
