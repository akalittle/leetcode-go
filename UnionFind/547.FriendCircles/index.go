package UnionFind

//There are N students in a class.Some of them are friends, while some are not.Their friendship is transitive in nature.For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C.And we defined a friend circle is a group of students who are direct or indirect friends.
//
//Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not.And you have to output the total number of friend circles among all the students.
//
//Example 1:
//Input:
//[[1, 1, 0],
//[1,1, 0],
//[0, 0,1]]
//Output: 2
//Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
//The 2nd student himself is in a friend circle. So return 2.
//Example 2:
//Input:
//[[1, 1, 0],
//[1, 1, 1],
//[0, 1, 1]]
//Output: 1
//Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
//so the 0th and 2nd students are indirect friends.All of them are in the same friend circle, so return 1.
//Note:
//N is in range [1, 200].
//M[i][i] = 1 for all students.
//If M[i][j] = 1, then M[j][i] = 1.

type UF struct {
	count  int
	parent []int
	rank   []int // rank[i]表示以i为根的集合中元素个数

}

func New(size int) *UF {
	rank := make([]int, size)
	parent := make([]int, size)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		rank[i] = 1
		parent[i] = i
	}

	return &UF{
		rank:   rank,
		parent: parent,
		count:  len(parent),
	}
}

func (uf *UF) find(p int) int {
	if p < 0 || p > len(uf.parent) {
		panic("p is out of range.")
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}

	return p
}

func (uf *UF) IsConnected(p int, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) UnionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}

	uf.count--

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上

	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = uf.parent[qRoot]
	} else if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = uf.parent[pRoot]
	} else { // rank[pRoot] == rank[qRoot]
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot] += 1 // 此时, 我维护rank的值
	}
}

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	length := len(M)
	uf := New(length)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if M[i][j] == 1 {
				uf.UnionElements(i, j)
			}
		}
	}
	return uf.count

}
