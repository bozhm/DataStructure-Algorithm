package main

import "fmt"

type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int
}

func (receiver *Stack) Push(num int) {
	if receiver.Top == receiver.MaxTop-1 {
		fmt.Println("满了")
		return
	}
	receiver.Top++
	receiver.arr[receiver.Top] = num
}

func (receiver *Stack) Show() {
	if receiver.Top == -1 {
		return
	}
	for i := receiver.Top; i >= 0; i-- {
		fmt.Println(receiver.arr[i])
	}
}

func (receiver *Stack) Pop() int {
	defer func() {
		receiver.Top--
	}()
	if receiver.Top == -1 {
		return -1
	}
	return receiver.arr[receiver.Top]
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}

	stack.Push(7)
	stack.Push(2)
	stack.Push(6)
	stack.Push(8)
	stack.Push(1)
	stack.Show()
	fmt.Println("-------------------------")
	stack.Pop()
	stack.Pop()
	stack.Show()

}
