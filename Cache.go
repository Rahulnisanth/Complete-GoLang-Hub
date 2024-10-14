package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, -num)
	}

	var result int64
	for i := 0; i < k; i++ {
		item := -heap.Pop(h).(int)
		heap.Push(h, -int(math.Ceil(float64(item)/3)))
		result += int64(item)
	}
	return result
}

func main() {
	nums := []int{1, 10, 3, 3, 3}
	k := 3
	fmt.Println(maxKelements(nums, k))
}
