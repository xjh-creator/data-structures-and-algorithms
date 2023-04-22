package stack

/*
	使用栈实现队列的下列操作：

	push(x) -- 将一个元素放入队列的尾部。
	pop() -- 从队列首部移除元素。
	peek() -- 返回队列首部的元素。
	empty() -- 返回队列是否为空。

	示例:

	MyQueue queue = new MyQueue();
	queue.push(1);
	queue.push(2);
	queue.peek();  // 返回 1
	queue.pop();   // 返回 1
	queue.empty(); // 返回 false

	说明:

    你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
    你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
    假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

*/

// 232.用栈实现队列
type MyQueue struct {
	stackIn  []int //输入栈
	stackOut []int //输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

// 往输入栈做push
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// 在输出栈做pop，pop时如果输出栈数据为空，需要将输入栈全部数据导入，如果非空，则可直接使用
func (this *MyQueue) Pop() int {
	inLen, outLen := len(this.stackIn), len(this.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}      //导出后清空
		outLen = len(this.stackOut) //更新长度值
	}
	val := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[:outLen-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

// isValid 20. 有效的括号
func isValid(s string) bool {
	hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}