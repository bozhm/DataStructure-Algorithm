package main

import (
	"errors"
	"fmt"
)

/**
数组队列
*/

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (receiver *Queue) Add(num int) (err error) {
	if receiver.rear == receiver.maxSize-1 {
		err = errors.New("队列已满")
		return
	}
	receiver.rear++
	receiver.array[receiver.rear] = num
	return
}

func (receiver *Queue) Show() {
	for i := receiver.front + 1; i <= receiver.rear; i++ {
		fmt.Printf("array[%v]=%v   ", i, receiver.array[i])
	}
}

func (receiver *Queue) Get() int {
	if receiver.front < receiver.rear {
		receiver.front++
		return receiver.array[receiver.front]
	}
	return -1

}

func main() {
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	for i := 0; i < 5; i++ {
		err := queue.Add(i)
		if err != nil {
			fmt.Println("当i=", i, err.Error())
		}
	}
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	queue.Show()

}
