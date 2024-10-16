package main

import (
	"container/heap"
	"math"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Smallest Chair :
func smallestChair(times [][]int, targetFriend int) int {
	chairs := make([]int, 0, len(times))
	for i := 0; i < len(times); i++ {
		chairs = append(chairs, i)
	}
	chairsHeap := IntHeap(chairs)
	heap.Init(&chairsHeap)

	leaveTimeToChairs := make(map[int][]int)
	leaveTimeHeap := IntHeap(make([]int, 0, len(times)))

	friends := make([]int, len(times))
	for i := 0; i < len(times); i++ {
		friends[i] = i
	}
	sort.Slice(friends, func(i, j int) bool {
		return times[friends[i]][0] < times[friends[j]][0]
	})

	for _, friend := range friends {
		curTime, leaveTime := times[friend][0], times[friend][1]

		for len(leaveTimeHeap) > 0 && leaveTimeHeap[0] <= curTime {
			top := heap.Pop(&leaveTimeHeap).(int)
			for _, chair := range leaveTimeToChairs[top] {
				heap.Push(&chairsHeap, chair)
			}
			delete(leaveTimeToChairs, top)
		}

		chair := heap.Pop(&chairsHeap).(int)
		if friend == targetFriend {
			return chair
		}
		heap.Push(&leaveTimeHeap, leaveTime)
		leaveTimeToChairs[leaveTime] = append(leaveTimeToChairs[leaveTime], chair)
	}
	return -1
}


// MAXIMUM SCORE AFTER APPLYING K OPERATIONS :
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
