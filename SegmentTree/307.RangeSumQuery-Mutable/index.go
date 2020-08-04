package SegmentTree

//Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
//
//The update(i, val) function modifies nums by updating the element at index i to val.
//
//Example:
//
//Given nums = [1, 3, 5]
//
//sumRange(0, 2) -> 9
//update(1, 2)
//sumRange(0, 2) -> 8
// 
//
//Constraints:
//
//The array is only modifiable by the update function.
//You may assume the number of calls to update and sumRange function is distributed evenly.
//0 <= i <= j <= nums.length - 1

type SegmentTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(interface{}, interface{}) interface{}
}

func New(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		tree:   make([]interface{}, len(arr)*4),
		data:   make([]interface{}, len(arr)),
		merger: merger,
	}

	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}

	if len(arr) == 0 {
		return segmentTree
	}

	segmentTree.buildSegmentTree(0, 0, len(arr)-1)

	return segmentTree
}

// 在treeIndex的位置创建表示区间[l...r]的线段树
func (st *SegmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("Index is illegal.")
	}

	return st.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

// 返回区间[queryL, queryR]的值
func (st *SegmentTree) Query(queryL int, queryR int) interface{} {
	if queryL < 0 || queryL >= len(st.data) || queryR < 0 || queryR > len(st.data) {
		panic("Index is illegal.")
	}

	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (st *SegmentTree) query(treeIndex int, l int, r int, queryL int, queryR int) interface{} {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	mid := l + (r-l)/2
	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	if queryL >= mid+1 {
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	leftResult := st.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)

	return st.merger(leftResult, rightResult)
}

func (st *SegmentTree) Set(index int, e interface{}) {
	if index < 0 || index > len(st.data) {
		panic("Index is Illegal")
	}

	st.data[index] = e
	st.set(0, 0, len(st.data)-1, index, e)
}

func (st *SegmentTree) set(treeIndex int, l int, r int, index int, e interface{}) {
	if l == r {
		st.tree[treeIndex] = e
		return
	}

	mid := l + (r-l)/2
	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	if index >= mid+1 {
		st.set(rightTreeIndex, mid+1, r, index, e)
	} else {
		st.set(leftTreeIndex, l, mid, index, e)
	}

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

type NumArray struct {
	segmentTree *SegmentTree
}

func Constructor(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	segTree := New(data, func(i interface{}, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	return NumArray{segTree}
}

func (na *NumArray) Update(i int, val int) {
	na.segmentTree.Set(i, val)
}

func (na *NumArray) SumRange(i int, j int) int {
	return na.segmentTree.Query(i, j).(int)
}
