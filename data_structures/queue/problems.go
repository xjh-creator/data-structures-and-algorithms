package queue
/*
	使用队列实现栈的下列操作：

    push(x) -- 元素 x 入栈
    pop() -- 移除栈顶元素
    top() -- 获取栈顶元素
    empty() -- 返回栈是否为空

注意:

    你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
    你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
    你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

*/
// 225. 用队列实现栈
type MyStack struct {
	//创建两个队列
	queue1 []int
	queue2 []int
}


func Constructor() MyStack {
	return MyStack{	//初始化
		queue1:make([]int,0),
		queue2:make([]int,0),
	}
}


func (this *MyStack) Push(x int)  {
	//先将数据存在queue2中
	this.queue2 = append(this.queue2,x)
	//将queue1中所有元素移到queue2中，再将两个队列进行交换
	this.Move()
}


func (this *MyStack) Move(){
	if len(this.queue1) == 0{
		//交换，queue1置为queue2,queue2置为空
		this.queue1,this.queue2 = this.queue2,this.queue1
	}else{
		//queue1元素从头开始一个一个追加到queue2中
		this.queue2 = append(this.queue2,this.queue1[0])
		this.queue1 = this.queue1[1:]	//去除第一个元素
		this.Move()     //重复
	}
}

func (this *MyStack) Pop() int {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]	//去除第一个元素
	return val

}


func (this *MyStack) Top() int {
	return this.queue1[0]	//直接返回
}


func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}