package queue

/*
	使用队列实现栈的下列操作：

    push(x) -- 元素 x 入栈
    pop() -- 移除栈顶元素
    top() -- 获取栈顶元素
    empty() -- 返回栈是否为空

*/

type MyStack struct {
	queue *ArrayQueue
}

func NewMyStack(n int) *MyStack {
	return &MyStack{
		queue: NewArrayQueue(n),
	}
}

func (m *MyStack)Push(e int) bool {
	return m.queue.EnQueue(e)
}

func (m *MyStack)Pop() int {
   return m.queue.DeQueue().(int)
}

func (m *MyStack)Top() int {
	return m.queue.Top().(int)
}

func (m *MyStack)Empty() bool {
	return m.queue.head == m.queue.tail
}

/*
	给定一个数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。

	你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。

	返回滑动窗口中的最大值。
*/
// MaxSlidingWindow 滑动窗口最大值
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}