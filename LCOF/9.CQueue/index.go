package LCOF

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 
//
//示例 1：
//
//输入：
//["CQueue", "appendTail", "deleteHead", "deleteHead"]
//[[],[3], [], []]
//输出：[null, null, 3,-1]
//示例 2：
//
//输入：
//["CQueue", "deleteHead", "appendTail", "appendTail", "deleteHead","deleteHead"]
//[[], [],[5], [2], [], []]
//输出：[null,-1, null, null, 5, 2]

type CQueue struct {
	in  []int // in 栈：存数据
	out []int // out 栈：出数据
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 直接入 int 栈
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	// 两个栈都为空
	if len(this.out) == 0 && len(this.in) == 0 {
		return -1
	}

	// out 栈为空，依次弹出 in 栈栈顶元素，入 out 栈
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			lastIndex := len(this.in) - 1
			popValue := this.in[lastIndex]
			this.in = this.in[:lastIndex]
			this.out = append(this.out, popValue)
		}
	}

	// 弹出 out 栈顶元素
	lastIndex := len(this.out) - 1
	popValue := this.out[lastIndex]
	this.out = this.out[:lastIndex]
	return popValue
}
