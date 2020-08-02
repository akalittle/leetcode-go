package Array

//Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
//
//Example:
//
//Given nums = [-2, 0, 3, -5, 2, -1]
//
//sumRange(0, 2) -> 1
//sumRange(2, 5) -> -1
//sumRange(0, 5) -> -3
// 
//
//Constraints:
//
//You may assume that the array does not change.
//There are many calls to sumRange function.
//0 <= nums.length <= 10^4
//-10^5 <= nums[i] <= 10^5
//0 <= i <= j < nums.length

type NumArray struct {
	tree   []int
	data   []int
	merger func(int, int) int
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

func (st *NumArray) buildSegmentTree(treeIndex int, l int, r int) {
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

func (st *NumArray) Query(queryL int, queryR int) int {
	if queryL < 0 || queryL >= len(st.data) || queryR < 0 || queryR > len(st.data) || queryL > queryR {
		//panic("Index is illegal.")
		return 0
	}

	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

func (st *NumArray) query(treeIndex int, l int, r int, queryL int, queryR int) int {
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

func Constructor(nums []int) *NumArray {
	segmentTree := &NumArray{
		tree: make([]int, len(nums)*4),
		data: make([]int, len(nums)),
		merger: func(i int, i2 int) int {
			return i + i2
		},
	}

	if len(nums) == 0 {
		return segmentTree
	}
	for i := 0; i < len(nums); i++ {
		segmentTree.data[i] = nums[i]
	}

	segmentTree.buildSegmentTree(0, 0, len(nums)-1)

	return segmentTree
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Query(i, j)
}
