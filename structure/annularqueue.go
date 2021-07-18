package main

import "fmt"

type AnnularQueue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (receiver *AnnularQueue) Add(num int) {
	if receiver.IsFull() {
		fmt.Println("满了加不了")
		return
	}
	receiver.array[receiver.rear] = num
	receiver.rear = (receiver.rear + 1) % receiver.maxSize
}

func (receiver *AnnularQueue) Get() int {
	if receiver.IsEmpty() {
		fmt.Println("队列为空")
		return -1
	}
	value := receiver.array[receiver.front]
	receiver.front = (receiver.front + 1) % receiver.maxSize
	return value
}

func (receiver *AnnularQueue) Show() {
	size := receiver.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	front := receiver.front
	for i := 0; i < size; i++ {
		fmt.Printf("array[%v]=%v   ", front, receiver.array[front])
		front = (receiver.front + 1) % receiver.maxSize
	}

}

// IsFull 判断是否满
func (receiver *AnnularQueue) IsFull() bool {
	return (receiver.rear+1)%receiver.maxSize == receiver.front
}

// IsEmpty 判断是否为空
func (receiver *AnnularQueue) IsEmpty() bool {
	return receiver.front == receiver.rear
}

// Size 获取size
func (receiver *AnnularQueue) Size() int {
	return (receiver.rear + receiver.maxSize - receiver.front) % receiver.maxSize
}

func main() {
	queue := &AnnularQueue{
		maxSize: 3,
		front:   0,
		rear:    0,
	}

	queue.Add(7)
	queue.Add(8)
	queue.Show()
	fmt.Println()
	queue.Get()
	queue.Get()
	queue.Show()
	fmt.Println()
	queue.Add(9)
	queue.Add(8)
	queue.Add(7)
	queue.Show()
}
