package DataStruct

import "fmt"

const initSize int = 20

type Stack struct {
	size int
	top  int
	data []interface{}
}

// NewStack 创建并初始化栈，返回strck
func NewStack() Stack {
	s := Stack{}
	s.size = initSize
	s.top = -1
	s.data = make([]interface{}, initSize)
	return s
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// IsFull 判断栈是否已满
func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

// Push 入栈
func (s *Stack) Push(data interface{}) bool {
	// 首先判断栈是否已满
	if s.IsFull() {
		fmt.Println("stack is full, push failed")
		return false
	}
	// 栈顶指针+1
	s.top++
	// 把当前的元素放在栈顶的位置
	s.data[s.top] = data
	return true
}

// Pop 返回栈顶元素
func (s *Stack) Pop() interface{} {
	// 判断是否是空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty , pop error")
		return -1
	}
	// 把栈顶的元素赋值给临时变量tmp
	tmp := s.data[s.top]
	// 栈顶指针-1
	s.top--
	return tmp
}

// GetLength 栈的元素的长度
func (s *Stack) GetLength() int {
	length := s.top + 1
	return length
}

// Clear 清空栈
func (s *Stack) Clear() {
	ss := NewStack()
	s = &(ss)
}

// Traverse 遍历栈
func (s *Stack) Traverse() {
	// 是否为空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}

	for i := 0; i <= s.top; i++ {
		fmt.Println(s.data[i], " ")
	}
}
