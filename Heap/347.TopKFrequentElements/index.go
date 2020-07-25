package Heap

import "container/heap"

//Given a non-empty array of integers, return the k most frequent elements.
//
//Example 1:
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//Example 2:
//
//Input: nums = [1], k = 1
//Output: [1]
//Note:
//
//You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
//Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
//It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
//You can return the answer in any order.

type PQ struct {
	val int
	cnt int
}

type PQMinHeap []PQ

func (pq *PQMinHeap) Len() int {
	return len(*pq)
}
func (pq *PQMinHeap) Less(i, j int) bool {
	return (*pq)[i].cnt < (*pq)[j].cnt
}
func (pq *PQMinHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PQMinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(PQ))
}

func (pq *PQMinHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}

func (pq *PQMinHeap) Peek() PQ {
	return (*pq)[0]
}

func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}
	minHeap := &PQMinHeap{}
	for key, val := range cntMap {
		heap.Push(minHeap, PQ{
			val: key,
			cnt: val,
		})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	res := make([]int, k)
	i := 0
	for minHeap.Len() > 0 {
		res[i] = minHeap.Pop().(PQ).val
		i++
	}
	return res
}
